package messages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lunaris/p10go/pkg/types"
)

type ChannelMode struct {
	OpMode                 bool
	Source                 types.ClientID
	Channel                string
	AddChannelModes        *types.ChannelModes
	RemoveChannelModes     *types.ChannelModes
	AddChannelUserModes    []types.ChannelMember
	RemoveChannelUserModes []types.ChannelMember
	Key                    string
	AddLimit               int
}

func ParseChannelMode(tokens []string) (*ChannelMode, error) {
	if len(tokens) < 4 {
		return nil, fmt.Errorf("MODE(CHANNEL): expected at least 4 tokens; received %d", len(tokens))
	}

	source, err := types.ParseClientID(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("MODE(CHANNEL): couldn't parse source: %w", err)
	}

	opMode := false
	if tokens[1] == "OM" {
		opMode = true
	} else if tokens[1] != "M" {
		return nil, fmt.Errorf("MODE(CHANNEL): expected M or OM; received %s", tokens[1])
	}

	channel := tokens[2]
	if channel[0] != '#' {
		return nil, fmt.Errorf("MODE(CHANNEL): expected channel to start with #; received %s", channel)
	}

	nextIndex := 4

	adding := true

	var addChannelModes *types.ChannelModes
	var removeChannelModes *types.ChannelModes

	key := ""
	addLimit := 0

	var addChannelUserModes []types.ChannelMember
	var removeChannelUserModes []types.ChannelMember

	seen := map[rune]bool{}

	for _, c := range tokens[3] {
		withChannelModes := func(f func(*types.ChannelModes)) {
			if adding {
				if addChannelModes == nil {
					addChannelModes = &types.ChannelModes{}
				}

				f(addChannelModes)
			} else {
				if removeChannelModes == nil {
					removeChannelModes = &types.ChannelModes{}
				}

				f(removeChannelModes)
			}
		}

		registerChannelMember := func(m types.ChannelMember) {
			if adding {
				addChannelUserModes = append(addChannelUserModes, m)
			} else {
				removeChannelUserModes = append(removeChannelUserModes, m)
			}
		}

		if c == '+' {
			adding = true
			continue
		}

		if c == '-' {
			adding = false
			continue
		}

		if c != 'o' && c != 'v' && seen[c] {
			return nil, fmt.Errorf("MODE(CHANNEL): duplicate mode: %c", c)
		}

		seen[c] = true

		switch c {
		case 'C':
			withChannelModes(func(m *types.ChannelModes) { m.NoCTCP = true })
		case 'c':
			withChannelModes(func(m *types.ChannelModes) { m.NoColour = true })
		case 'D':
			withChannelModes(func(m *types.ChannelModes) { m.DelayedJoins = true })
		case 'i':
			withChannelModes(func(m *types.ChannelModes) { m.InviteOnly = true })
		case 'k':
			withChannelModes(func(m *types.ChannelModes) { m.Keyed = true })
			key = tokens[nextIndex]
			nextIndex++
		case 'l':
			withChannelModes(func(m *types.ChannelModes) { m.Limit = true })
			if adding {
				limit, err := strconv.Atoi(tokens[nextIndex])
				if err != nil {
					return nil, fmt.Errorf("MODE(CHANNEL): couldn't parse channel limit: %w", err)
				}

				addLimit = limit
				nextIndex++
			}
		case 'M':
			withChannelModes(func(m *types.ChannelModes) { m.UnregisteredModerated = true })
		case 'm':
			withChannelModes(func(m *types.ChannelModes) { m.Moderated = true })
		case 'N':
			withChannelModes(func(m *types.ChannelModes) { m.NoNotices = true })
		case 'n':
			withChannelModes(func(m *types.ChannelModes) {
				m.NoPrivateMessages = true
			})
		case 'o':
			clientId, err := types.ParseClientID(tokens[nextIndex])
			if err != nil {
				return nil, fmt.Errorf("MODE(CHANNEL): couldn't parse client ID: %w", err)
			}

			registerChannelMember(types.ChannelMember{
				ClientID: clientId,
				Modes:    types.ChannelUserModes{Op: true},
			})

			nextIndex++
		case 'p':
			withChannelModes(func(m *types.ChannelModes) { m.Private = true })
		case 'r':
			withChannelModes(func(m *types.ChannelModes) { m.RegisteredOnly = true })
		case 's':
			withChannelModes(func(m *types.ChannelModes) { m.Secret = true })
		case 'T':
			withChannelModes(func(m *types.ChannelModes) { m.NoMultipleTargets = true })
		case 't':
			withChannelModes(func(m *types.ChannelModes) { m.TopicLimited = true })
		case 'u':
			withChannelModes(func(m *types.ChannelModes) { m.NoQuitParts = true })
		case 'v':
			clientId, err := types.ParseClientID(tokens[nextIndex])
			if err != nil {
				return nil, fmt.Errorf("MODE(CHANNEL): couldn't parse client ID: %w", err)
			}

			registerChannelMember(types.ChannelMember{
				ClientID: clientId,
				Modes:    types.ChannelUserModes{Voice: true},
			})

			nextIndex++
		}
	}

	return &ChannelMode{
		OpMode:                 opMode,
		Source:                 source,
		Channel:                channel,
		AddChannelModes:        addChannelModes,
		RemoveChannelModes:     removeChannelModes,
		AddChannelUserModes:    addChannelUserModes,
		RemoveChannelUserModes: removeChannelUserModes,
		Key:                    key,
		AddLimit:               addLimit,
	}, nil
}

func (m *ChannelMode) String() string {
	var modes strings.Builder
	var args strings.Builder

	token := "M"
	if m.OpMode {
		token = "OM"
	}

	if m.AddChannelModes != nil {
		s := m.AddChannelModes.String()
		if len(s) > 0 {
			modes.WriteString("+" + s)
		}

		if m.AddChannelModes.Keyed && m.Key != "" {
			args.WriteString(" " + m.Key)
		}

		if m.AddChannelModes.Limit && m.AddLimit > 0 {
			args.WriteString(" " + strconv.Itoa(m.AddLimit))
		}
	}

	if m.RemoveChannelModes != nil {
		s := m.RemoveChannelModes.String()
		if len(s) > 0 {
			modes.WriteString("-" + s)
		}

		if m.RemoveChannelModes.Keyed && m.Key != "" {
			args.WriteString(" " + m.Key)
		}
	}

	for _, m := range m.AddChannelUserModes {
		if m.Modes.Op {
			modes.WriteString("+o")
			args.WriteString(" " + m.ClientID.String())
		}
		if m.Modes.Voice {
			modes.WriteString("+v")
			args.WriteString(" " + m.ClientID.String())
		}
	}

	for _, m := range m.RemoveChannelUserModes {
		if m.Modes.Op {
			modes.WriteString("-o")
			args.WriteString(" " + m.ClientID.String())
		}
		if m.Modes.Voice {
			modes.WriteString("-v")
			args.WriteString(" " + m.ClientID.String())
		}
	}

	return fmt.Sprintf(
		"%s %s %s %s%s",
		m.Source,
		token,
		m.Channel,
		modes.String(),
		args.String(),
	)
}
