package client

import (
	"github.com/lunaris/p10go/pkg/messages"
)

func (c *P10Client) handleEndOfBurst(m *messages.EndOfBurst) {
	c.debugf("received END_OF_BURST; sending acknowledgement", "numeric", m.ServerNumeric)
	c.Send(&messages.EndOfBurstAcknowledgement{ServerNumeric: c.config.ClientNumeric})
}
