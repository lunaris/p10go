package operserv

import (
	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/logging"
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

type Operserv struct {
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

func NewOperserv(config Configuration) *Operserv {
	return &Operserv{
		logger: config.Logger,

		clientID: config.ClientID,
		nick:     config.Nick,
		info:     config.Info,
		maskUser: config.MaskUser,
		maskHost: config.MaskHost,
	}
}

func (o *Operserv) OnEvent(cl *client.P10Client, e client.Event) {
	switch e.Type {
	case client.MessageEvent:
		switch m := e.Message.(type) {
		case *messages.EndOfBurst:
			o.handleEndOfBurst(cl, m)
		}
	}
}
