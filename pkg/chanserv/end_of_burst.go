package chanserv

import (
	"time"

	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

func (c *Chanserv) handleEndOfBurst(cl *client.P10Client, m *messages.EndOfBurst) {
	c.infof("sending nick", "info", c.info, "maskUser", c.maskUser, "maskHost", c.maskHost)
	cl.Send(&messages.Nick{
		ServerNumeric:    c.clientID.Server,
		Nick:             c.nick,
		HopCount:         1,
		CreatedTimestamp: time.Now().Unix(),
		MaskUser:         c.maskUser,
		MaskHost:         c.maskHost,
		UserModes:        types.UserModes{Invisible: true},
		IP:               "+6n",
		ClientID:         c.clientID,
		Info:             c.info,
	})
	cl.Send(&messages.Join{
		ClientID:  c.clientID,
		Channel:   "#dev",
		Timestamp: time.Now().Unix(),
	})
	cl.Send(&messages.ChannelMode{
		OpMode:  true,
		Source:  c.clientID,
		Channel: "#dev",
		AddChannelUserModes: []types.ChannelMember{
			{
				ClientID: c.clientID,
				Modes:    types.ChannelUserModes{Op: true},
			},
		},
	})
}
