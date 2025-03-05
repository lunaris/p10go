package client

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/lunaris/p10go/pkg/logging"
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

type P10Client struct {
	context context.Context
	logger  logging.Logger

	serverAddress string

	clientPassword    string
	clientNumeric     types.ServerNumeric
	clientName        string
	clientDescription string

	conn     net.Conn
	buf      []byte
	servers  map[types.ServerNumeric]*Server
	clients  map[types.ClientID]*Client
	channels map[string]*Channel

	observers []Observer

	done chan struct{}
}

type Configuration struct {
	Context context.Context
	Logger  logging.Logger

	ServerAddress string

	ClientPassword    string
	ClientNumeric     types.ServerNumeric
	ClientName        string
	ClientDescription string

	Observers []Observer
}

type Observer interface {
	OnEvent(*P10Client, Event)
}

type Server struct {
	Numeric types.ServerNumeric
}

type Client struct {
	ID   types.ClientID
	Nick string
}

type EventType string

const (
	MessageEvent EventType = "message"
	ErrorEvent   EventType = "error"
	CloseEvent   EventType = "close"
)

type Event struct {
	Type    EventType
	Message messages.Message
	Error   error
}

func Connect(config Configuration) (*P10Client, error) {
	conn, err := net.Dial("tcp", config.ServerAddress)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to server: %w", err)
	}

	buf := make([]byte, 1024)
	servers := make(map[types.ServerNumeric]*Server)
	clients := make(map[types.ClientID]*Client)
	channels := make(map[string]*Channel)

	c := &P10Client{
		context: config.Context,
		logger:  config.Logger,

		serverAddress: config.ServerAddress,

		clientPassword:    config.ClientPassword,
		clientNumeric:     config.ClientNumeric,
		clientName:        config.ClientName,
		clientDescription: config.ClientDescription,

		conn:     conn,
		buf:      buf,
		servers:  servers,
		clients:  clients,
		channels: channels,

		observers: config.Observers,

		done: make(chan struct{}),
	}

	c.debugf("connected to server")
	go c.eventLoop()

	return c, nil
}

func (c *P10Client) Close() {
	c.debugf("closing connection")
	c.notifyObservers(Event{Type: CloseEvent})
	c.conn.Close()
	close(c.done)
}

func (c *P10Client) Done() <-chan struct{} {
	return c.done
}

func (c *P10Client) eventLoop() {
	defer c.Close()

	err := c.handshake()
	if err != nil {
		c.notifyObservers(Event{Type: ErrorEvent, Error: err})
		return
	}

	for {
		ms, err := c.receive()
		if err != nil {
			c.notifyObservers(Event{Type: ErrorEvent, Error: err})
			return
		}

		for _, m := range ms {
			switch m := m.(type) {
			case *messages.EndOfBurst:
				c.handleEndOfBurst(m)
			case *messages.Join:
				c.handleJoin(m)
			case *messages.Nick:
				c.handleNick(m)
			case *messages.Ping:
				c.handlePing(m)
			case *messages.Server:
				c.handleServer(m)
			}

			c.notifyObservers(Event{Type: MessageEvent, Message: m})
		}
	}
}

func (c *P10Client) handshake() error {
	err := c.Send(&messages.Pass{Password: c.clientPassword})
	if err != nil {
		return fmt.Errorf("couldn't send PASS: %w", err)
	}

	err = c.Send(&messages.Server{
		Name:           c.clientName,
		HopCount:       1,
		StartTimestamp: time.Now().Unix(),
		LinkTimestamp:  time.Now().Unix(),
		Protocol:       messages.J10,
		Numeric:        c.clientNumeric,
		MaxConnections: "]]]",
		Description:    c.clientDescription,
	})
	if err != nil {
		return fmt.Errorf("couldn't send SERVER: %w", err)
	}

	return nil
}

func (c *P10Client) Send(m messages.Message) error {
	c.debugf("sending message", "message", m.String())

	_, err := c.conn.Write([]byte(m.String() + "\r\n"))
	if err != nil {
		return fmt.Errorf("couldn't send message: %w", err)
	}

	return nil
}

var lineBreak = regexp.MustCompile(`\r?\n`)

func (c *P10Client) receive() ([]messages.Message, error) {
	n, err := c.conn.Read(c.buf)
	if err != nil {
		return nil, fmt.Errorf("couldn't receive message: %w", err)
	}

	bs := c.buf[:n]
	c.debugf("received bytes", "bytes", string(bs))

	lines := lineBreak.Split(string(bs), -1)

	ms := make([]messages.Message, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		tokens := strings.Split(line, " ")

		c.debugf("parsing tokens", "lineNumber", i, "tokens", tokens)
		m, err := messages.Parse(tokens)
		if err != nil {
			c.errorf("couldn't parse message", "line", line, "error", err)
			return nil, fmt.Errorf("couldn't parse message: %w", err)
		}

		c.debugf("parsed message", "lineNumber", i, "message", m.String())
		ms[i] = m
	}

	return ms, nil
}

func (c *P10Client) notifyObservers(e Event) {
	for _, o := range c.observers {
		o.OnEvent(c, e)
	}
}

func (c *P10Client) Channel(name string) *Channel {
	ch, ok := c.channels[name]
	if !ok {
		ch = NewChannel(name)
		c.channels[name] = ch
	}

	return ch
}
