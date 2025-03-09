package messages

import (
	"fmt"
	"strings"

	"github.com/lunaris/p10go/pkg/types"
)

type UserMode struct {
	OpMode      bool
	Source      types.ClientID
	Nick        string
	AddModes    *types.UserModes
	RemoveModes *types.UserModes
}

func ParseUserMode(tokens []string) (*UserMode, error) {
	if len(tokens) < 4 {
		return nil, fmt.Errorf("MODE(USER): expected at least 4 tokens; received %d", len(tokens))
	}

	source, err := types.ParseClientID(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("MODE(USER): couldn't parse source: %w", err)
	}

	opMode := false
	if tokens[1] == "OM" {
		opMode = true
	} else if tokens[1] != "M" {
		return nil, fmt.Errorf("MODE(USER): expected M or OM; received %s", tokens[1])
	}

	nick := tokens[2]

	adding := true

	var addModes *types.UserModes
	var removeModes *types.UserModes

	seen := map[rune]bool{}

	for _, c := range tokens[3] {
		withModes := func(f func(*types.UserModes)) {
			if adding {
				if addModes == nil {
					addModes = &types.UserModes{}
				}

				f(addModes)
			} else {
				if removeModes == nil {
					removeModes = &types.UserModes{}
				}

				f(removeModes)
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

		if seen[c] {
			return nil, fmt.Errorf("MODE(USER): duplicate mode: %c", c)
		}

		seen[c] = true

		switch c {
		case 'd':
			withModes(func(m *types.UserModes) { m.Deaf = true })
		case 'g':
			withModes(func(m *types.UserModes) { m.Debug = true })
		case 'h':
			withModes(func(m *types.UserModes) { m.SetHost = true })
		case 'I':
			withModes(func(m *types.UserModes) { m.NoIdle = true })
		case 'i':
			withModes(func(m *types.UserModes) { m.Invisible = true })
		case 'k':
			withModes(func(m *types.UserModes) { m.ChannelService = true })
		case 'n':
			withModes(func(m *types.UserModes) { m.NoChannels = true })
		case 'O':
			withModes(func(m *types.UserModes) { m.LocalOp = true })
		case 'o':
			withModes(func(m *types.UserModes) { m.Op = true })
		case 'P':
			withModes(func(m *types.UserModes) { m.Paranoid = true })
		case 'R':
			withModes(func(m *types.UserModes) { m.AccountOnly = true })
		case 'r':
			withModes(func(m *types.UserModes) { m.Account = true })
		case 's':
			withModes(func(m *types.UserModes) { m.ServerNotices = true })
		case 'w':
			withModes(func(m *types.UserModes) { m.WallOps = true })
		case 'X':
			withModes(func(m *types.UserModes) { m.ExtraOp = true })
		case 'x':
			withModes(func(m *types.UserModes) { m.HiddenHost = true })
		}
	}

	return &UserMode{
		OpMode:      opMode,
		Source:      source,
		Nick:        nick,
		AddModes:    addModes,
		RemoveModes: removeModes,
	}, nil
}

func (m *UserMode) String() string {
	var modes strings.Builder

	token := "M"
	if m.OpMode {
		token = "OM"
	}

	if m.AddModes != nil {
		s := m.AddModes.String()
		if len(s) > 0 {
			modes.WriteString("+" + s)
		}
	}

	if m.RemoveModes != nil {
		s := m.RemoveModes.String()
		if len(s) > 0 {
			modes.WriteString("-" + s)
		}
	}

	return fmt.Sprintf(
		"%s %s %s %s",
		m.Source,
		token,
		m.Nick,
		modes.String(),
	)
}
