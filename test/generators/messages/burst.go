package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedBan = rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,15}$`)

var GeneratedBurst = rapid.Custom(func(t *rapid.T) *messages.Burst {
	channelModes := typeGenerators.GeneratedChannelModes.Draw(t, "ChannelModes")
	channelLimit := 0
	if channelModes.Limit {
		channelLimit = rapid.IntRange(1, 100).Draw(t, "ChannelLimit")
	}

	channelKey := ""
	if channelModes.Keyed {
		channelKey = rapid.StringMatching(`^[A-Za-z0-9-]{1,32}$`).Draw(t, "ChannelKey")
	}

	return &messages.Burst{
		ServerNumeric:    typeGenerators.GeneratedServerNumeric.Draw(t, "ServerNumeric"),
		Channel:          rapid.StringMatching(`^[#][A-Za-z][A-Za-z0-9-]{1,31}$`).Draw(t, "Channel"),
		CreatedTimestamp: rapid.Int64().Draw(t, "CreatedTimestamp"),
		ChannelModes:     channelModes,
		ChannelLimit:     channelLimit,
		ChannelKey:       channelKey,
		Members:          rapid.SliceOfN(typeGenerators.GeneratedChannelMember, 1, 10).Draw(t, "Members"),
		Bans:             rapid.SliceOfN(GeneratedBan, 0, 10).Draw(t, "Bans"),
	}
})
