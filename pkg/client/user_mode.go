package client

import (
	"github.com/lunaris/p10go/pkg/messages"
)

func (c *P10Client) handleUserMode(m *messages.UserMode) {
	c.debugf("received MODE/OPMODE; updating user", "source", m.Source, "nick", m.Nick)

	u := c.usersByNick[m.Nick]
	if u == nil {
		c.debugf("received MODE/OPMODE for unknown user", "nick", m.Nick)
		return
	}

	if m.AddModes != nil {
		u.modes.Add(*m.AddModes)
	}
	if m.RemoveModes != nil {
		u.modes.Remove(*m.RemoveModes)
	}
}
