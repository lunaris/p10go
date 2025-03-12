package client

import (
	"github.com/lunaris/p10go/pkg/messages"
)

func (c *P10Client) handleNick(m *messages.Nick) {
	c.debugf("saw NICK; updating users", "id", m.ClientID, "nick", m.Nick)

	u := c.addUser(m.ClientID, m.Nick)
	*u.modes = m.UserModes
}
