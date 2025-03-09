package messages

import (
	"fmt"

	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/pkg/types"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedChannelMode = rapid.Custom(func(t *rapid.T) *messages.ChannelMode {
	channelMode := &messages.ChannelMode{
		OpMode:  rapid.Bool().Draw(t, "OpMode"),
		Source:  typeGenerators.GeneratedClientID.Draw(t, "Source"),
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
			ClientID: typeGenerators.GeneratedClientID.Draw(t, fmt.Sprintf("AddChannelUserModes[%d].ClientID", i)),
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
			ClientID: typeGenerators.GeneratedClientID.Draw(t, fmt.Sprintf("RemoveChannelUserModes[%d].ClientID", i)),
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
