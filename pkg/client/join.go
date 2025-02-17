package client

import (
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

func (c *P10Client) handleJoin(m *messages.Join) {
	c.debugf("received JOIN; updating channels", "client", m.ClientID, "channel", m.Channel)

	client := c.clients[m.ClientID]
	if client == nil {
		c.debugf("received JOIN for unknown client", "client", m.ClientID)
		return
	}

	ch := c.Channel(m.Channel)
	ch.members[m.ClientID] = &types.ChannelMember{ClientID: m.ClientID}
}
