package client

import (
	"github.com/lunaris/p10go/pkg/messages"
)

func (c *P10Client) handleBurst(m *messages.Burst) {
	c.debugf("received BURST; updating channels", "channel", m.Channel)

	ch := c.channel(m.Channel)

	*ch.modes = m.ChannelModes
	ch.limit = m.ChannelLimit
	ch.key = m.ChannelKey

	for _, member := range m.Members {
		chm := member
		ch.members[member.ClientID] = &chm
	}
}
