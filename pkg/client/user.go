package client

import "github.com/lunaris/p10go/pkg/types"

type user struct {
	id    types.ClientID
	nick  string
	modes *types.UserModes
}

func (c *P10Client) addUser(id types.ClientID, nick string) *user {
	u := &user{
		id:    id,
		nick:  nick,
		modes: &types.UserModes{},
	}

	c.usersByClientID[id] = u
	c.usersByNick[nick] = u

	return u
}
