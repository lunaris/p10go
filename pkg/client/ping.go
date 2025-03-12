package client

import (
	"github.com/lunaris/p10go/pkg/messages"
)

func (c *P10Client) handlePing(m *messages.Ping) {
	c.debugf("saw PING; sending PONG", "source", m.Source)
	c.Send(&messages.Pong{Source: "QQ", Target: m.Source})
}
