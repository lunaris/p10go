package client

import "github.com/lunaris/p10go/pkg/types"

type Channel struct {
	name    string
	members map[types.ClientID]*types.ChannelMember
}

func NewChannel(name string) *Channel {
	return &Channel{
		name:    name,
		members: make(map[types.ClientID]*types.ChannelMember),
	}
}
