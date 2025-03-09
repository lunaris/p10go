package generators

import (
	"github.com/lunaris/p10go/pkg/types"
	"pgregory.net/rapid"
)

var GeneratedServerNumeric = rapid.Custom(func(t *rapid.T) types.ServerNumeric {
	return types.ServerNumeric(
		rapid.StringMatching(`^[A-Za-z0-9\[\]]{2}`).Draw(t, "ServerNumeric"),
	)
})

var GeneratedClientNumeric = rapid.Custom(func(t *rapid.T) types.ClientNumeric {
	return types.ClientNumeric(
		rapid.StringMatching(`^[A-Za-z0-9\[\]]{3}`).Draw(t, "ClientNumeric"),
	)
})

var GeneratedClientID = rapid.Custom(func(t *rapid.T) types.ClientID {
	server := GeneratedServerNumeric.Draw(t, "Server")
	client := GeneratedClientNumeric.Draw(t, "Client")

	return types.ClientID{
		Server: server,
		Client: client,
	}
})

var GeneratedChannelModes = rapid.Custom(func(t *rapid.T) types.ChannelModes {
	return types.ChannelModes{
		NoCTCP:                rapid.Bool().Draw(t, "NoCTCP"),
		NoColour:              rapid.Bool().Draw(t, "NoColour"),
		DelayedJoins:          rapid.Bool().Draw(t, "DelayedJoins"),
		InviteOnly:            rapid.Bool().Draw(t, "InviteOnly"),
		Keyed:                 rapid.Bool().Draw(t, "Keyed"),
		Limit:                 rapid.Bool().Draw(t, "Limit"),
		UnregisteredModerated: rapid.Bool().Draw(t, "UnregisteredModerated"),
		Moderated:             rapid.Bool().Draw(t, "Moderated"),
		NoNotices:             rapid.Bool().Draw(t, "NoNotices"),
		NoPrivateMessages:     rapid.Bool().Draw(t, "NoPrivateMessages"),
		Private:               rapid.Bool().Draw(t, "Private"),
		RegisteredOnly:        rapid.Bool().Draw(t, "RegisteredOnly"),
		Secret:                rapid.Bool().Draw(t, "Secret"),
		NoMultipleTargets:     rapid.Bool().Draw(t, "NoMultipleTargets"),
		TopicLimited:          rapid.Bool().Draw(t, "TopicLimited"),
		NoQuitParts:           rapid.Bool().Draw(t, "NoQuitParts"),
	}
})

var GeneratedUserModes = rapid.Custom(func(t *rapid.T) types.UserModes {
	return types.UserModes{
		Deaf:           rapid.Bool().Draw(t, "Deaf"),
		Debug:          rapid.Bool().Draw(t, "Debug"),
		SetHost:        rapid.Bool().Draw(t, "SetHost"),
		NoIdle:         rapid.Bool().Draw(t, "NoIdle"),
		Invisible:      rapid.Bool().Draw(t, "Invisible"),
		ChannelService: rapid.Bool().Draw(t, "ChannelService"),
		NoChannels:     rapid.Bool().Draw(t, "NoChannels"),
		LocalOp:        rapid.Bool().Draw(t, "LocalOp"),
		Op:             rapid.Bool().Draw(t, "Op"),
		Paranoid:       rapid.Bool().Draw(t, "Paranoid"),
		AccountOnly:    rapid.Bool().Draw(t, "AccountOnly"),
		Account:        rapid.Bool().Draw(t, "Account"),
		ServerNotices:  rapid.Bool().Draw(t, "ServerNotices"),
		WallOps:        rapid.Bool().Draw(t, "WallOps"),
		ExtraOp:        rapid.Bool().Draw(t, "ExtraOp"),
		HiddenHost:     rapid.Bool().Draw(t, "HiddenHost"),
	}
})

var GeneratedChannelUserModes = rapid.Custom(func(t *rapid.T) types.ChannelUserModes {
	return types.ChannelUserModes{
		Op:    rapid.Bool().Draw(t, "Op"),
		Voice: rapid.Bool().Draw(t, "Voice"),
	}
})

var GeneratedChannelMember = rapid.Custom(func(t *rapid.T) types.ChannelMember {
	return types.ChannelMember{
		ClientID: GeneratedClientID.Draw(t, "ClientID"),
		Modes:    GeneratedChannelUserModes.Draw(t, "Modes"),
	}
})
