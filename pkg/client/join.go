package client

import (
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

func (c *P10Client) handleJoin(m *messages.Join) {
	c.debugf("received JOIN; updating channels", "client", m.ClientID, "channel", m.Channel)

	u := c.usersByClientID[m.ClientID]
	if u == nil {
		c.debugf("received JOIN for unknown user", "client", m.ClientID)
		return
	}

	ch := c.channel(m.Channel)
	ch.members[m.ClientID] = &types.ChannelMember{ClientID: m.ClientID}
}
