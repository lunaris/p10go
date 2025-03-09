package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedUserMode = rapid.Custom(func(t *rapid.T) *messages.UserMode {
	userMode := &messages.UserMode{
		OpMode: rapid.Bool().Draw(t, "OpMode"),
		Source: typeGenerators.GeneratedClientID.Draw(t, "Source"),
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
