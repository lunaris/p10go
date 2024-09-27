package client

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/lunaris/p10go/messages"
	"github.com/lunaris/p10go/types"
)

type P10Client struct {
	config   Configuration
	conn     net.Conn
	buf      []byte
	servers  map[types.ServerNumeric]*Server
	clients  map[types.ClientID]*Client
	channels map[string]*Channel
	events   chan Event
}

type Configuration struct {
	Context context.Context
	Logger  *slog.Logger

	ServerAddress string

	ClientPassword    string
	ClientNumeric     types.ServerNumeric
	ClientName        string
	ClientDescription string
}

type Server struct {
	Numeric types.ServerNumeric
}

type Client struct {
	ID   types.ClientID
	Nick string
}

type Channel struct{}

type EventType string

const (
	MessageEvent EventType = "message"
	ErrorEvent   EventType = "error"
)

type Event struct {
	Type    EventType
	Message messages.Message
	Error   error
}

func New(config Configuration) (*P10Client, error) {
	conn, err := net.Dial("tcp", config.ServerAddress)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to server: %w", err)
	}

	buf := make([]byte, 1024)
	servers := make(map[types.ServerNumeric]*Server)
	clients := make(map[types.ClientID]*Client)
	channels := make(map[string]*Channel)
	events := make(chan Event)

	c := &P10Client{
		config:   config,
		conn:     conn,
		buf:      buf,
		servers:  servers,
		clients:  clients,
		channels: channels,
		events:   events,
	}

	c.debug("connected to server")
	go c.eventLoop()

	return c, nil
}

func (c *P10Client) Close() {
	c.debug("closing connection")
	close(c.events)
	c.conn.Close()
}

func (c *P10Client) Events() <-chan Event {
	return c.events
}

func (c *P10Client) eventLoop() {
	err := c.handshake()
	if err != nil {
		c.events <- Event{Type: ErrorEvent, Error: err}
		c.Close()
		return
	}

	for {
		ms, err := c.receive()
		if err != nil {
			c.events <- Event{Type: ErrorEvent, Error: err}
			c.Close()
			return
		}

		for _, m := range ms {
			switch m := m.(type) {
			case *messages.EndOfBurst:
				c.debug("received END_OF_BURST; sending acknowledgement", "numeric", m.ServerNumeric)
				c.Send(&messages.EndOfBurstAcknowledgement{ServerNumeric: "QQ"})
			case *messages.Nick:
				c.debug("received NICK; updating clients", "id", m.ClientID, "nick", m.Nick)
				c.clients[m.ClientID] = &Client{
					ID:   m.ClientID,
					Nick: m.Nick,
				}
			case *messages.Ping:
				c.debug("received PING; sending PONG", "source", m.Source)
				c.Send(&messages.Pong{Source: "QQ", Target: m.Source})
			case *messages.Server:
				c.debug("received SERVER; updating servers", "numeric", m.Numeric)
				c.servers[m.Numeric] = &Server{
					Numeric: m.Numeric,
				}
			}

			c.events <- Event{Type: MessageEvent, Message: m}
		}
	}
}

func (c *P10Client) handshake() error {
	err := c.Send(&messages.Pass{Password: c.config.ClientPassword})
	if err != nil {
		return fmt.Errorf("couldn't send PASS: %w", err)
	}

	err = c.Send(&messages.Server{
		Name:           c.config.ClientName,
		HopCount:       1,
		StartTimestamp: time.Now().Unix(),
		LinkTimestamp:  time.Now().Unix(),
		Protocol:       messages.J10,
		Numeric:        c.config.ClientNumeric,
		MaxConnections: "]]]",
		Description:    c.config.ClientDescription,
	})
	if err != nil {
		return fmt.Errorf("couldn't send SERVER: %w", err)
	}

	return nil
}

func (c *P10Client) Send(m messages.Message) error {
	c.debug("sending message", "message", m.String())

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
	c.debug("received bytes", "bytes", string(bs))

	lines := lineBreak.Split(string(bs), -1)

	ms := make([]messages.Message, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		tokens := strings.Split(line, " ")

		c.debug("parsing tokens", "index", i, "tokens", tokens)
		m, err := messages.Parse(tokens)
		if err != nil {
			c.error("couldn't parse message", "line", line, "error", err)
			return nil, fmt.Errorf("couldn't parse message: %w", err)
		}

		c.info("parsed message", "index", i, "message", m.String())
		ms[i] = m
	}

	return ms, nil
}

func (c *P10Client) log(level slog.Level, message string, args ...interface{}) {
	if c.config.Logger != nil {
		c.config.Logger.Log(c.config.Context, level, message, args...)
	}
}

func (c *P10Client) debug(message string, args ...interface{}) {
	c.log(slog.LevelDebug, message, args...)
}

func (c *P10Client) info(message string, args ...interface{}) {
	c.log(slog.LevelInfo, message, args...)
}

func (c *P10Client) warn(message string, args ...interface{}) {
	c.log(slog.LevelWarn, message, args...)
}

func (c *P10Client) error(message string, args ...interface{}) {
	c.log(slog.LevelError, message, args...)
}
