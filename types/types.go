package types

import (
	"fmt"
	"regexp"
	"strings"
)

type ServerNumeric string

var serverNumericPattern = regexp.MustCompile(`^[A-Za-z0-9\[\]]{2}$`)

func ParseServerNumeric(s string) (ServerNumeric, error) {
	if !serverNumericPattern.MatchString(s) {
		return "", fmt.Errorf("invalid server numeric: %s", s)
	}

	return ServerNumeric(s), nil
}

type ClientNumeric string

var clientNumericPattern = regexp.MustCompile(`^[A-Za-z0-9\[\]]{3}$`)

func ParseClientNumeric(s string) (ClientNumeric, error) {
	if !clientNumericPattern.MatchString(s) {
		return "", fmt.Errorf("invalid client numeric: %s", s)
	}

	return ClientNumeric(s), nil
}

type ClientID struct {
	Server ServerNumeric
	Client ClientNumeric
}

func ParseClientID(s string) (ClientID, error) {
	if len(s) != 5 {
		return ClientID{}, fmt.Errorf("invalid client ID: %s", s)
	}

	server, err := ParseServerNumeric(s[:2])
	if err != nil {
		return ClientID{}, fmt.Errorf("couldn't parse server numeric: %w", err)
	}

	client, err := ParseClientNumeric(s[2:])
	if err != nil {
		return ClientID{}, fmt.Errorf("couldn't parse client numeric: %w", err)
	}

	return ClientID{
		Server: server,
		Client: client,
	}, nil
}

func (id ClientID) String() string {
	return fmt.Sprintf("%s%s", id.Server, id.Client)
}

type ChannelModes struct {
	// C
	NoCTCP bool
	// c
	NoColour bool
	// D
	DelayedJoins bool
	// i
	InviteOnly bool
	// k
	Keyed bool
	// l
	Limit bool
	// M
	UnregisteredModerated bool
	// m
	Moderated bool
	// N
	NoNotices bool
	// n
	NoPrivateMessages bool
	// p
	Private bool
	// r
	RegisteredOnly bool
	// s
	Secret bool
	// T
	NoMultipleTargets bool
	// t
	TopicLimited bool
	// u
	NoQuitParts bool
}

func ParseChannelModes(s string) (*ChannelModes, error) {
	modes := &ChannelModes{}
	var invalidModes []rune

	for _, r := range s {
		switch r {
		case 'C':
			modes.NoCTCP = true
		case 'c':
			modes.NoColour = true
		case 'D':
			modes.DelayedJoins = true
		case 'i':
			modes.InviteOnly = true
		case 'k':
			modes.Keyed = true
		case 'l':
			modes.Limit = true
		case 'M':
			modes.UnregisteredModerated = true
		case 'm':
			modes.Moderated = true
		case 'N':
			modes.NoNotices = true
		case 'n':
			modes.NoPrivateMessages = true
		case 'p':
			modes.Private = true
		case 'r':
			modes.RegisteredOnly = true
		case 's':
			modes.Secret = true
		case 'T':
			modes.NoMultipleTargets = true
		case 't':
			modes.TopicLimited = true
		case 'u':
			modes.NoQuitParts = true
		default:
			invalidModes = append(invalidModes, r)
		}
	}

	if len(invalidModes) > 0 {
		return nil, fmt.Errorf("invalid channel modes: %s", string(invalidModes))
	}

	return modes, nil
}

func (m *ChannelModes) String() string {
	var sb strings.Builder

	if m.NoCTCP {
		sb.WriteRune('C')
	}
	if m.NoColour {
		sb.WriteRune('c')
	}
	if m.DelayedJoins {
		sb.WriteRune('D')
	}
	if m.InviteOnly {
		sb.WriteRune('i')
	}
	if m.Keyed {
		sb.WriteRune('k')
	}
	if m.Limit {
		sb.WriteRune('l')
	}
	if m.UnregisteredModerated {
		sb.WriteRune('M')
	}
	if m.Moderated {
		sb.WriteRune('m')
	}
	if m.NoNotices {
		sb.WriteRune('N')
	}
	if m.NoPrivateMessages {
		sb.WriteRune('n')
	}
	if m.Private {
		sb.WriteRune('p')
	}
	if m.RegisteredOnly {
		sb.WriteRune('r')
	}
	if m.Secret {
		sb.WriteRune('s')
	}
	if m.NoMultipleTargets {
		sb.WriteRune('T')
	}
	if m.TopicLimited {
		sb.WriteRune('t')
	}
	if m.NoQuitParts {
		sb.WriteRune('u')
	}

	return sb.String()
}

type UserModes struct {
	// d
	Deaf bool
	// g
	Debug bool
	// h
	SetHost bool
	// I
	NoIdle bool
	// i
	Invisible bool
	// k
	ChannelService bool
	// n
	NoChannels bool
	// O
	LocalOp bool
	// o
	Op bool
	// P
	Paranoid bool
	// R
	AccountOnly bool
	// r
	Account bool
	// s
	ServerNotices bool
	// w
	WallOps bool
	// X
	ExtraOp bool
	// x
	HiddenHost bool
}

func ParseUserModes(s string) (*UserModes, error) {
	modes := &UserModes{}
	var invalidModes []rune

	for _, r := range s {
		switch r {
		case 'd':
			modes.Deaf = true
		case 'g':
			modes.Debug = true
		case 'h':
			modes.SetHost = true
		case 'I':
			modes.NoIdle = true
		case 'i':
			modes.Invisible = true
		case 'k':
			modes.ChannelService = true
		case 'n':
			modes.NoChannels = true
		case 'O':
			modes.LocalOp = true
		case 'o':
			modes.Op = true
		case 'P':
			modes.Paranoid = true
		case 'R':
			modes.AccountOnly = true
		case 'r':
			modes.Account = true
		case 's':
			modes.ServerNotices = true
		case 'w':
			modes.WallOps = true
		case 'X':
			modes.ExtraOp = true
		case 'x':
			modes.HiddenHost = true
		default:
			invalidModes = append(invalidModes, r)
		}
	}

	if len(invalidModes) > 0 {
		return nil, fmt.Errorf("invalid user modes: %s", string(invalidModes))
	}

	return modes, nil
}

func (m *UserModes) String() string {
	var sb strings.Builder

	if m.Deaf {
		sb.WriteRune('d')
	}
	if m.Debug {
		sb.WriteRune('g')
	}
	if m.SetHost {
		sb.WriteRune('h')
	}
	if m.NoIdle {
		sb.WriteRune('I')
	}
	if m.Invisible {
		sb.WriteRune('i')
	}
	if m.ChannelService {
		sb.WriteRune('k')
	}
	if m.NoChannels {
		sb.WriteRune('n')
	}
	if m.LocalOp {
		sb.WriteRune('O')
	}
	if m.Op {
		sb.WriteRune('o')
	}
	if m.Paranoid {
		sb.WriteRune('P')
	}
	if m.AccountOnly {
		sb.WriteRune('R')
	}
	if m.Account {
		sb.WriteRune('r')
	}
	if m.ServerNotices {
		sb.WriteRune('s')
	}
	if m.WallOps {
		sb.WriteRune('w')
	}
	if m.ExtraOp {
		sb.WriteRune('X')
	}
	if m.HiddenHost {
		sb.WriteRune('x')
	}

	return sb.String()
}

type ChannelUserModes struct {
	// o
	Op bool
	// v
	Voice bool
}

func ParseChannelUserModes(s string) (*ChannelUserModes, error) {
	modes := &ChannelUserModes{}
	var invalidModes []rune

	for _, r := range s {
		switch r {
		case 'o':
			modes.Op = true
		case 'v':
			modes.Voice = true
		default:
			invalidModes = append(invalidModes, r)
		}
	}

	if len(invalidModes) > 0 {
		return nil, fmt.Errorf("invalid channel user modes: %s", string(invalidModes))
	}

	return modes, nil
}

func (m *ChannelUserModes) String() string {
	var sb strings.Builder

	if m.Op {
		sb.WriteRune('o')
	}
	if m.Voice {
		sb.WriteRune('v')
	}

	return sb.String()
}

type ChannelMember struct {
	ClientID ClientID
	Modes    *ChannelUserModes
}

func ParseChannelMember(s string) (*ChannelMember, error) {
	parts := strings.Split(s, ":")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid channel client: %s", s)
	}

	clientID, err := ParseClientID(parts[0])
	if err != nil {
		return nil, fmt.Errorf("couldn't parse client ID: %w", err)
	}

	if len(parts) == 1 {
		return &ChannelMember{ClientID: clientID}, nil
	}

	modes, err := ParseChannelUserModes(parts[1])
	if err != nil {
		return nil, fmt.Errorf("couldn't parse channel user modes: %w", err)
	}

	return &ChannelMember{ClientID: clientID, Modes: modes}, nil
}

func (m *ChannelMember) String() string {
	if m.Modes == nil {
		return m.ClientID.String()
	}

	return fmt.Sprintf("%s:%s", m.ClientID, m.Modes)
}

func ParseChannelMembers(s string) ([]*ChannelMember, error) {
	parts := strings.Split(s, ",")
	members := make([]*ChannelMember, len(parts))
	for i, part := range parts {
		member, err := ParseChannelMember(part)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse channel user: %w", err)
		}

		members[i] = member
	}

	return members, nil
}
