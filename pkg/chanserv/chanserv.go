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

	clientID types.ClientID
	nick     string
	info     string
	maskUser string
	maskHost string
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

		clientID: config.ClientID,
		nick:     config.Nick,
		info:     config.Info,
		maskUser: config.MaskUser,
		maskHost: config.MaskHost,
	}
}

func (c *Chanserv) OnEvent(cl *client.P10Client, e client.Event) {
	switch e.Type {
	case client.MessageEvent:
		switch m := e.Message.(type) {
		case *messages.EndOfBurst:
			c.handleEndOfBurst(cl, m)
		case *messages.Join:
			c.handleJoin(cl, m)
		}
	}
}

func (c *Chanserv) handleEndOfBurst(cl *client.P10Client, m *messages.EndOfBurst) {
	c.infof("sending nick", "info", c.info, "maskUser", c.maskUser, "maskHost", c.maskHost)
	cl.Send(&messages.Nick{
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
	cl.Send(&messages.Join{
		ClientID:  c.clientID,
		Channel:   "#dev",
		Timestamp: time.Now().Unix(),
	})
	cl.Send(&messages.ChannelMode{
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

func (c *Chanserv) handleJoin(cl *client.P10Client, m *messages.Join) {
	cl.Send(&messages.ChannelMode{
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
