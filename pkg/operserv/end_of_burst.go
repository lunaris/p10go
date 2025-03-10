package operserv

import (
	"time"

	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

func (o *Operserv) handleEndOfBurst(cl *client.P10Client, m *messages.EndOfBurst) {
	o.infof("sending nick", "info", o.info, "maskUser", o.maskUser, "maskHost", o.maskHost)
	cl.Send(&messages.Nick{
		ServerNumeric:    o.clientID.Server,
		Nick:             o.nick,
		HopCount:         1,
		CreatedTimestamp: time.Now().Unix(),
		MaskUser:         o.maskUser,
		MaskHost:         o.maskHost,
		UserModes:        types.UserModes{Invisible: true, Op: true},
		IP:               "+6n",
		ClientID:         o.clientID,
		Info:             o.info,
	})
	cl.Send(&messages.Join{
		ClientID:  o.clientID,
		Channel:   "#dev",
		Timestamp: time.Now().Unix(),
	})
	cl.Send(&messages.ChannelMode{
		OpMode:  true,
		Source:  o.clientID,
		Channel: "#dev",
		AddChannelUserModes: []types.ChannelMember{
			{
				ClientID: o.clientID,
				Modes:    types.ChannelUserModes{Op: true},
			},
		},
	})
}
