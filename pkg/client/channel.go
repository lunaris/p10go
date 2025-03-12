package client

import "github.com/lunaris/p10go/pkg/types"

type channel struct {
	name    string
	modes   *types.ChannelModes
	limit   int
	key     string
	members map[types.ClientID]*types.ChannelMember
}

func (c *P10Client) channel(name string) *channel {
	if ch, ok := c.channels[name]; ok {
		return ch
	}

	ch := &channel{
		name:    name,
		modes:   &types.ChannelModes{},
		members: make(map[types.ClientID]*types.ChannelMember),
	}

	c.channels[name] = ch
	return ch
}
