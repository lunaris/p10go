package client

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"regexp"
	"strings"

	"github.com/lunaris/p10go/messages"
	"github.com/lunaris/p10go/types"
)

type P10Client struct {
	ctx      context.Context
	logger   *slog.Logger
	conn     net.Conn
	buf      []byte
	servers  map[types.ServerNumeric]*Server
	clients  map[types.ClientID]*Client
	channels map[string]*Channel
}

type Server struct {
	Numeric types.ServerNumeric
}

type Client struct {
	ID   types.ClientID
	Nick string
}

type Channel struct{}

func New(ctx context.Context, address string, opts *Options) (*P10Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to server: %w", err)
	}

	buf := make([]byte, 1024)
	servers := make(map[types.ServerNumeric]*Server)
	clients := make(map[types.ClientID]*Client)
	channels := make(map[string]*Channel)

	c := &P10Client{
		ctx:      ctx,
		conn:     conn,
		buf:      buf,
		servers:  servers,
		clients:  clients,
		channels: channels,
	}

	if opts != nil {
		c.logger = opts.Logger
	}

	c.debug("connected to server")

	return c, nil
}

type Options struct {
	Logger *slog.Logger
}

func (c *P10Client) Close() {
	c.debug("closing connection")
	c.conn.Close()
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

func (c *P10Client) Receive() ([]messages.Message, error) {
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
	}

	return ms, nil
}

func (c *P10Client) log(level slog.Level, message string, args ...interface{}) {
	if c.logger != nil {
		c.logger.Log(c.ctx, level, message, args...)
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
