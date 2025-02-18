package chanserv

import (
	"time"

	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/logging"
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

type Chanserv struct {
	logger logging.Logger
	client *client.P10Client

	clientID types.ClientID
	nick     string
	info     string
	maskUser string
	maskHost string

	done chan struct{}
}

type Configuration struct {
	Logger logging.Logger
	Client *client.P10Client

	ClientID types.ClientID
	Nick     string
	Info     string
	MaskUser string
	MaskHost string
}

func NewChanserv(config Configuration) *Chanserv {
	return &Chanserv{
		logger: config.Logger,
		client: config.Client,

		clientID: config.ClientID,
		nick:     config.Nick,
		info:     config.Info,
		maskUser: config.MaskUser,
		maskHost: config.MaskHost,

		done: make(chan struct{}),
	}
}

func (c *Chanserv) Go() {
	for e := range c.client.Events() {
		switch e.Type {
		case client.MessageEvent:
			switch m := e.Message.(type) {
			case *messages.EndOfBurst:
				c.handleEndOfBurst(m)
			case *messages.Join:
				c.handleJoin(m)
			}
		}
	}

	close(c.done)
}

func (c *Chanserv) Done() <-chan struct{} {
	return c.done
}

func (c *Chanserv) handleEndOfBurst(m *messages.EndOfBurst) {
	c.infof("sending nick", "nick", c.nick, "info", c.info, "maskUser", c.maskUser, "maskHost", c.maskHost)
	c.client.Send(&messages.Nick{
		ServerNumeric:    c.clientID.Server,
		Nick:             c.nick,
		HopCount:         1,
		CreatedTimestamp: time.Now().Unix(),
		MaskUser:         c.maskUser,
		MaskHost:         c.maskHost,
		UserModes:        types.UserModes{Invisible: true},
		IP:               "+6n",
		ClientID:         c.clientID,
		Info:             c.info,
	})
	c.client.Send(&messages.Join{
		ClientID:  c.clientID,
		Channel:   "#dev",
		Timestamp: time.Now().Unix(),
	})
	c.client.Send(&messages.ChannelMode{
		OpMode:  true,
		Source:  c.clientID,
		Channel: "#dev",
		AddChannelUserModes: []types.ChannelMember{
			{
				ClientID: c.clientID,
				Modes:    types.ChannelUserModes{Op: true},
			},
		},
	})
}

func (c *Chanserv) handleJoin(m *messages.Join) {
	c.client.Send(&messages.ChannelMode{
		OpMode:  true,
		Source:  c.clientID,
		Channel: m.Channel,
		AddChannelUserModes: []types.ChannelMember{
			{
				ClientID: m.ClientID,
				Modes:    types.ChannelUserModes{Op: true},
			},
		},
	})
}
