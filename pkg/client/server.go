package client

import (
	"github.com/lunaris/p10go/pkg/messages"
)

func (c *P10Client) handleServer(m *messages.Server) {
	c.debugf("received SERVER; updating servers", "numeric", m.Numeric)
	c.servers[m.Numeric] = &Server{
		Numeric: m.Numeric,
	}
}
