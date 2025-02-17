package chanserv

import (
	"time"

	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/logging"
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

type Chanserv struct {
	config Configuration

	done chan struct{}
}

type Configuration struct {
	Logger logging.Logger
	Client *client.P10Client
}

func NewChanserv(config Configuration) *Chanserv {
	return &Chanserv{
		config: config,
		done:   make(chan struct{}),
	}
}

func (c *Chanserv) Go() {
	for e := range c.config.Client.Events() {
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
	c.infof("sending Q", "m", m)
	c.config.Client.Send(&messages.Nick{
		ServerNumeric:    "QQ",
		Nick:             "Q",
		HopCount:         1,
		CreatedTimestamp: time.Now().Unix(),
		MaskUser:         "Q",
		MaskHost:         "services.p10.localhost",
		UserModes:        types.UserModes{Invisible: true},
		IP:               "+6n",
		ClientID: types.ClientID{
			Server: "QQ",
			Client: "AAA",
		},
		Info: "The Q Bot",
	})
	c.config.Client.Send(&messages.Join{
		ClientID: types.ClientID{
			Server: "QQ",
			Client: "AAA",
		},
		Channel:   "#dev",
		Timestamp: time.Now().Unix(),
	})
	c.config.Client.Send(&messages.ChannelMode{
		OpMode: true,
		Source: types.ClientID{
			Server: "QQ",
			Client: "AAA",
		},
		Channel: "#dev",
		AddChannelUserModes: []types.ChannelMember{
			{
				ClientID: types.ClientID{
					Server: "QQ",
					Client: "AAA",
				},
				Modes: types.ChannelUserModes{Op: true},
			},
		},
	})
}

func (c *Chanserv) handleJoin(m *messages.Join) {
	c.config.Client.Send(&messages.ChannelMode{
		OpMode: true,
		Source: types.ClientID{
			Server: "QQ",
			Client: "AAA",
		},
		Channel: m.Channel,
		AddChannelUserModes: []types.ChannelMember{
			{
				ClientID: m.ClientID,
				Modes:    types.ChannelUserModes{Op: true},
			},
		},
	})
}
