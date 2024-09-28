package generators

import (
	"github.com/lunaris/p10go/pkg/messages"
	"pgregory.net/rapid"
)

var GeneratedPass = rapid.Custom(func(t *rapid.T) *messages.Pass {
	return &messages.Pass{
		Password: rapid.StringMatching(`^[A-Za-z][A-Za-z0-9-]{1,31}$`).Draw(t, "Password"),
	}
})

var GeneratedServer = rapid.Custom(func(t *rapid.T) *messages.Server {
	return &messages.Server{
		Name:           rapid.StringMatching(`^[A-Za-z][A-Za-z0-9-]{1,31}$`).Draw(t, "Name"),
		HopCount:       rapid.IntRange(1, 10).Draw(t, "HopCount"),
		StartTimestamp: rapid.Int64().Draw(t, "StartTimestamp"),
		LinkTimestamp:  rapid.Int64().Draw(t, "LinkTimestamp"),
		Protocol: rapid.SampledFrom([]messages.Protocol{
			messages.J10,
			messages.J10,
		}).Draw(t, "Protocol"),
		Numeric:        GeneratedServerNumeric.Draw(t, "Numeric"),
		MaxConnections: GeneratedClientNumeric.Draw(t, "MaxConnections"),
		Description:    rapid.StringMatching(`^[A-Za-z0-9- ]{0,64}$`).Draw(t, "Description"),
	}
})

var GeneratedNick = rapid.Custom(func(t *rapid.T) *messages.Nick {
	userModes := GeneratedUserModes.Draw(t, "UserModes")
	account := ""
	if userModes.Account {
		account = rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,15}$`).Draw(t, "Account")
	}

	return &messages.Nick{
		ServerNumeric:    GeneratedServerNumeric.Draw(t, "ServerNumeric"),
		Nick:             rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,19}$`).Draw(t, "Nick"),
		HopCount:         rapid.IntRange(1, 10).Draw(t, "HopCount"),
		CreatedTimestamp: rapid.Int64().Draw(t, "CreatedTimestamp"),
		MaskUser:         rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,15}$`).Draw(t, "MaskUser"),
		MaskHost:         rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,15}$`).Draw(t, "MaskHost"),
		UserModes:        userModes,
		Account:          account,
		IP:               rapid.StringMatching(`^[A-Za-z0-9\[\]]{3,5}$`).Draw(t, "IP"),
		ClientID:         GeneratedClientID.Draw(t, "ClientID"),
		Info:             rapid.StringMatching(`^[A-Za-z0-9- ]{0,64}$`).Draw(t, "Info"),
	}
})

var GeneratedBurst = rapid.Custom(func(t *rapid.T) *messages.Burst {
	channelModes := GeneratedChannelModes.Draw(t, "ChannelModes")
	channelLimit := 0
	if channelModes.Limit {
		channelLimit = rapid.IntRange(1, 100).Draw(t, "ChannelLimit")
	}

	channelKey := ""
	if channelModes.Keyed {
		channelKey = rapid.StringMatching(`^[A-Za-z0-9-]{1,32}$`).Draw(t, "ChannelKey")
	}

	return &messages.Burst{
		ServerNumeric:    GeneratedServerNumeric.Draw(t, "ServerNumeric"),
		Channel:          rapid.StringMatching(`^[#][A-Za-z][A-Za-z0-9-]{1,31}$`).Draw(t, "Channel"),
		CreatedTimestamp: rapid.Int64().Draw(t, "CreatedTimestamp"),
		ChannelModes:     channelModes,
		ChannelLimit:     channelLimit,
		ChannelKey:       channelKey,
		Members:          rapid.SliceOfN(GeneratedChannelMember, 1, 10).Draw(t, "Members"),
		Bans:             rapid.SliceOfN(GeneratedBan, 0, 10).Draw(t, "Bans"),
	}
})

var GeneratedBan = rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,15}$`)
