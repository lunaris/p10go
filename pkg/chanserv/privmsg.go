package chanserv

import (
	"strings"
	"time"

	"github.com/lunaris/p10go/pkg/chanserv/persistence"
	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/messages"
)

func (c *Chanserv) handlePrivmsg(cl *client.P10Client, m *messages.Privmsg) {
	parts := strings.Split(m.Message, " ")
	switch parts[0] {
	case "AUTH":
		if len(parts) != 3 {
			cl.Send(&messages.Privmsg{
				Source:  c.clientID,
				Target:  m.Source,
				Message: "Invalid AUTH command. Usage: AUTH <username> <password>",
			})

			return
		}

		username := parts[1]
		password := parts[2]

		_, err := c.users.Authenticate(persistence.AuthenticateUserRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			cl.Send(&messages.Privmsg{
				Source:  c.clientID,
				Target:  m.Source,
				Message: "Invalid AUTH credentials",
			})

			return
		}

		cl.Send(&messages.Account{
			Source:      c.clientID.Server,
			Target:      m.Source,
			AccountName: username,
			Timestamp:   time.Now().Unix(),
		})
	case "DEBUG":
		cl.Send(&rawMessage{raw: strings.Join(parts[1:], " ")})
	}
}

type rawMessage struct {
	raw string
}

func (r *rawMessage) String() string {
	return r.raw
}
