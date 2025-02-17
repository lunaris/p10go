package generators

import (
	"fmt"

	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
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

var GeneratedEndOfBurst = rapid.Custom(func(t *rapid.T) *messages.EndOfBurst {
	return &messages.EndOfBurst{
		ServerNumeric: GeneratedServerNumeric.Draw(t, "ServerNumeric"),
	}
})

var GeneratedEndOfBurstAcknowledgement = rapid.Custom(func(t *rapid.T) *messages.EndOfBurstAcknowledgement {
	return &messages.EndOfBurstAcknowledgement{
		ServerNumeric: GeneratedServerNumeric.Draw(t, "ServerNumeric"),
	}
})

var GeneratedPing = rapid.Custom(func(t *rapid.T) *messages.Ping {
	return &messages.Ping{
		Source: GeneratedServerNumeric.Draw(t, "Source"),
	}
})

var GeneratedPong = rapid.Custom(func(t *rapid.T) *messages.Pong {
	return &messages.Pong{
		Source: GeneratedServerNumeric.Draw(t, "Source"),
		Target: GeneratedServerNumeric.Draw(t, "Target"),
	}
})

var GeneratedChannelName = rapid.StringMatching(`^[#][A-Za-z][A-Za-z0-9-]{1,31}$`)

var GeneratedJoin = rapid.Custom(func(t *rapid.T) *messages.Join {
	return &messages.Join{
		ClientID:  GeneratedClientID.Draw(t, "ClientID"),
		Channel:   GeneratedChannelName.Draw(t, "Channel"),
		Timestamp: rapid.Int64().Draw(t, "Timestamp"),
	}
})

var GeneratedChannelMode = rapid.Custom(func(t *rapid.T) *messages.ChannelMode {
	channelMode := &messages.ChannelMode{
		OpMode:  rapid.Bool().Draw(t, "OpMode"),
		Source:  GeneratedClientID.Draw(t, "Source"),
		Channel: GeneratedChannelName.Draw(t, "Channel"),
	}

	doNothing := 0
	add := 1
	// remove := 2

	withAction := func(action int, f func(*types.ChannelModes)) {
		if action == doNothing {
			return
		}

		if action == add {
			if channelMode.AddChannelModes == nil {
				channelMode.AddChannelModes = &types.ChannelModes{}
			}

			f(channelMode.AddChannelModes)
		} else {
			if channelMode.RemoveChannelModes == nil {
				channelMode.RemoveChannelModes = &types.ChannelModes{}
			}

			f(channelMode.RemoveChannelModes)
		}
	}

	noCTCP := rapid.IntRange(0, 2).Draw(t, "NoCTCP")
	withAction(noCTCP, func(m *types.ChannelModes) { m.NoCTCP = true })

	noColour := rapid.IntRange(0, 2).Draw(t, "NoColour")
	withAction(noColour, func(m *types.ChannelModes) { m.NoColour = true })

	delayedJoins := rapid.IntRange(0, 2).Draw(t, "DelayedJoins")
	withAction(delayedJoins, func(m *types.ChannelModes) { m.DelayedJoins = true })

	inviteOnly := rapid.IntRange(0, 2).Draw(t, "InviteOnly")
	withAction(inviteOnly, func(m *types.ChannelModes) { m.InviteOnly = true })

	keyed := rapid.IntRange(0, 2).Draw(t, "Keyed")
	withAction(keyed, func(m *types.ChannelModes) { m.Keyed = true })

	if keyed != doNothing {
		channelMode.Key = rapid.StringMatching(`^[A-Za-z0-9-]{1,32}$`).Draw(t, "Key")
	}

	limit := rapid.IntRange(0, 2).Draw(t, "Limit")
	withAction(limit, func(m *types.ChannelModes) { m.Limit = true })
	if limit == add {
		channelMode.AddLimit = rapid.IntRange(1, 100).Draw(t, "Limit")
	}

	unregisteredModerated := rapid.IntRange(0, 2).Draw(t, "UnregisteredModerated")
	withAction(unregisteredModerated, func(m *types.ChannelModes) { m.UnregisteredModerated = true })

	moderated := rapid.IntRange(0, 2).Draw(t, "Moderated")
	withAction(moderated, func(m *types.ChannelModes) { m.Moderated = true })

	noNotices := rapid.IntRange(0, 2).Draw(t, "NoNotices")
	withAction(noNotices, func(m *types.ChannelModes) { m.NoNotices = true })

	noPrivateMessages := rapid.IntRange(0, 2).Draw(t, "NoPrivateMessages")
	withAction(noPrivateMessages, func(m *types.ChannelModes) { m.NoPrivateMessages = true })

	private := rapid.IntRange(0, 2).Draw(t, "Private")
	withAction(private, func(m *types.ChannelModes) { m.Private = true })

	registeredOnly := rapid.IntRange(0, 2).Draw(t, "RegisteredOnly")
	withAction(registeredOnly, func(m *types.ChannelModes) { m.RegisteredOnly = true })

	secret := rapid.IntRange(0, 2).Draw(t, "Secret")
	withAction(secret, func(m *types.ChannelModes) { m.Secret = true })

	noMultipleTargets := rapid.IntRange(0, 2).Draw(t, "NoMultipleTargets")
	withAction(noMultipleTargets, func(m *types.ChannelModes) { m.NoMultipleTargets = true })

	topicLimited := rapid.IntRange(0, 2).Draw(t, "TopicLimited")
	withAction(topicLimited, func(m *types.ChannelModes) { m.TopicLimited = true })

	noQuitParts := rapid.IntRange(0, 2).Draw(t, "NoQuitParts")
	withAction(noQuitParts, func(m *types.ChannelModes) { m.NoQuitParts = true })

	addChannelUserModesCount := rapid.IntRange(0, 3).Draw(t, "AddChannelUserModesCount")
	for i := 0; i < addChannelUserModesCount; i++ {
		channelMember := types.ChannelMember{
			ClientID: GeneratedClientID.Draw(t, fmt.Sprintf("AddChannelUserModes[%d].ClientID", i)),
		}

		shouldOp := rapid.Bool().Draw(t, fmt.Sprintf("AddChannelUserModes[%d].ShouldOp", i))
		if shouldOp {
			channelMember.Modes.Op = true
		} else {
			channelMember.Modes.Voice = true
		}

		channelMode.AddChannelUserModes = append(channelMode.AddChannelUserModes, channelMember)
	}

	removeChannelUserModesCount := rapid.IntRange(0, 3).Draw(t, "RemoveChannelUserModesCount")
	for i := 0; i < removeChannelUserModesCount; i++ {
		channelMember := types.ChannelMember{
			ClientID: GeneratedClientID.Draw(t, fmt.Sprintf("RemoveChannelUserModes[%d].ClientID", i)),
		}

		shouldOp := rapid.Bool().Draw(t, fmt.Sprintf("RemoveChannelUserModes[%d].ShouldOp", i))
		if shouldOp {
			channelMember.Modes.Op = true
		} else {
			channelMember.Modes.Voice = true
		}

		channelMode.AddChannelUserModes = append(channelMode.AddChannelUserModes, channelMember)
	}

	return channelMode
})

var GeneratedUserMode = rapid.Custom(func(t *rapid.T) *messages.UserMode {
	userMode := &messages.UserMode{
		OpMode: rapid.Bool().Draw(t, "OpMode"),
		Source: GeneratedClientID.Draw(t, "Source"),
		Nick:   rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,19}$`).Draw(t, "Nick"),
	}

	doNothing := 0
	add := 1
	// remove := 2

	withAction := func(action int, f func(*types.UserModes)) {
		if action == doNothing {
			return
		}

		if action == add {
			if userMode.AddModes == nil {
				userMode.AddModes = &types.UserModes{}
			}

			f(userMode.AddModes)
		} else {
			if userMode.RemoveModes == nil {
				userMode.RemoveModes = &types.UserModes{}
			}

			f(userMode.RemoveModes)
		}
	}

	deaf := rapid.IntRange(0, 2).Draw(t, "Deaf")
	withAction(deaf, func(m *types.UserModes) { m.Deaf = true })

	debug := rapid.IntRange(0, 2).Draw(t, "Debug")
	withAction(debug, func(m *types.UserModes) { m.Debug = true })

	setHost := rapid.IntRange(0, 2).Draw(t, "SetHost")
	withAction(setHost, func(m *types.UserModes) { m.SetHost = true })

	noIdle := rapid.IntRange(0, 2).Draw(t, "NoIdle")
	withAction(noIdle, func(m *types.UserModes) { m.NoIdle = true })

	invisible := rapid.IntRange(0, 2).Draw(t, "Invisible")
	withAction(invisible, func(m *types.UserModes) { m.Invisible = true })

	channelService := rapid.IntRange(0, 2).Draw(t, "ChannelService")
	withAction(channelService, func(m *types.UserModes) { m.ChannelService = true })

	noChannels := rapid.IntRange(0, 2).Draw(t, "NoChannels")
	withAction(noChannels, func(m *types.UserModes) { m.NoChannels = true })

	localOp := rapid.IntRange(0, 2).Draw(t, "LocalOp")
	withAction(localOp, func(m *types.UserModes) { m.LocalOp = true })

	op := rapid.IntRange(0, 2).Draw(t, "Op")
	withAction(op, func(m *types.UserModes) { m.Op = true })

	paranoid := rapid.IntRange(0, 2).Draw(t, "Paranoid")
	withAction(paranoid, func(m *types.UserModes) { m.Paranoid = true })

	accountOnly := rapid.IntRange(0, 2).Draw(t, "AccountOnly")
	withAction(accountOnly, func(m *types.UserModes) { m.AccountOnly = true })

	account := rapid.IntRange(0, 2).Draw(t, "Account")
	withAction(account, func(m *types.UserModes) { m.Account = true })

	serverNotices := rapid.IntRange(0, 2).Draw(t, "ServerNotices")
	withAction(serverNotices, func(m *types.UserModes) { m.ServerNotices = true })

	wallOps := rapid.IntRange(0, 2).Draw(t, "WallOps")
	withAction(wallOps, func(m *types.UserModes) { m.WallOps = true })

	extraOp := rapid.IntRange(0, 2).Draw(t, "ExtraOp")
	withAction(extraOp, func(m *types.UserModes) { m.ExtraOp = true })

	hiddenHost := rapid.IntRange(0, 2).Draw(t, "HiddenHost")
	withAction(hiddenHost, func(m *types.UserModes) { m.HiddenHost = true })

	return userMode
})
