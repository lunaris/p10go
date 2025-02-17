package client

import (
	"github.com/lunaris/p10go/pkg/messages"
)

func (c *P10Client) handleNick(m *messages.Nick) {
	c.debugf("received NICK; updating clients", "id", m.ClientID, "nick", m.Nick)
	c.clients[m.ClientID] = &Client{
		ID:   m.ClientID,
		Nick: m.Nick,
	}
}
