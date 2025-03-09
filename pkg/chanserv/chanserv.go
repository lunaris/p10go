package chanserv

import (
	"github.com/lunaris/p10go/pkg/chanserv/persistence"
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

	users persistence.UserRepository
}

type Configuration struct {
	Logger logging.Logger
	Client *client.P10Client

	ClientID types.ClientID
	Nick     string
	Info     string
	MaskUser string
	MaskHost string

	Users persistence.UserRepository
}

func NewChanserv(config Configuration) *Chanserv {
	return &Chanserv{
		logger: config.Logger,

		clientID: config.ClientID,
		nick:     config.Nick,
		info:     config.Info,
		maskUser: config.MaskUser,
		maskHost: config.MaskHost,

		users: config.Users,
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
		case *messages.Privmsg:
			c.handlePrivmsg(cl, m)
		}
	}
}
