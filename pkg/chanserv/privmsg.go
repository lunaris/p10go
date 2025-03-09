package chanserv

import (
	"strings"

	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
)

func (c *Chanserv) handlePrivmsg(cl *client.P10Client, m *messages.Privmsg) {
	parts := strings.Split(m.Message, " ")
	switch parts[0] {
	case "AUTH":
		cl.Send(&messages.UserMode{
			OpMode:   true,
			Source:   c.clientID,
			Nick:     m.Source.String(),
			AddModes: &types.UserModes{Account: true},
		})
	}
}
