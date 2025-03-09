package chanserv

import (
	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

func (c *Chanserv) handleJoin(cl *client.P10Client, m *messages.Join) {
	cl.Send(&messages.ChannelMode{
		OpMode:  true,
		Source:  c.clientID,
		Channel: m.Channel,
		AddChannelUserModes: []types.ChannelMember{
			{
				ClientID: m.ClientID,
				Modes:    types.ChannelUserModes{Op: true},
			},
		},
	})
}
