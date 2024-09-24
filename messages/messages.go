package messages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lunaris/p10go/types"
)

type Message interface {
	String() string
}

func Parse(tokens []string) (Message, error) {
	if len(tokens) < 2 {
		return nil, fmt.Errorf("expected at least 2 tokens; received %d", len(tokens))
	}

	switch tokens[0] {
	case "PASS":
		return ParsePass(tokens)
	case "SERVER":
		return ParseServer(tokens)
	}

	switch tokens[1] {
	case "B":
		return ParseBurst(tokens)
	case "EA":
		return ParseEndOfBurstAcknowledgement(tokens)
	case "EB":
		return ParseEndOfBurst(tokens)
	case "G":
		return ParsePing(tokens)
	case "N":
		return ParseNick(tokens)
	case "Z":
		return ParsePong(tokens)
	}

	return &Unknown{Tokens: tokens}, nil
}

type Pass struct {
	Password string
}

func ParsePass(tokens []string) (*Pass, error) {
	if len(tokens) != 2 {
		return nil, fmt.Errorf("PASS: expected 2 tokens; recevied %d", len(tokens))
	}

	if tokens[0] != "PASS" {
		return nil, fmt.Errorf("PASS: expected PASS; received %s", tokens[0])
	}

	return &Pass{Password: lastParameter(tokens[1:])}, nil
}

func (p *Pass) String() string {
	return fmt.Sprintf("PASS :%s", p.Password)
}

type Server struct {
	Name           string
	HopCount       int
	StartTimestamp int64
	LinkTimestamp  int64
	Protocol       Protocol
	Numeric        types.ServerNumeric
	MaxConnections types.ClientNumeric
	Description    string
}

func ParseServer(tokens []string) (*Server, error) {
	if len(tokens) < 8 {
		return nil, fmt.Errorf("SERVER: expected at least 8 tokens; received %d", len(tokens))
	}

	if tokens[0] != "SERVER" {
		return nil, fmt.Errorf("SERVER: expected SERVER; received %s", tokens[0])
	}

	hopCount, err := strconv.Atoi(tokens[2])
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse hop count: %w", err)
	}

	startTimestamp, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse start timestamp: %w", err)
	}

	linkTimestamp, err := strconv.ParseInt(tokens[4], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse link timestamp: %w", err)
	}

	protocol, err := ParseProtocol(tokens[5])
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse protocol: %w", err)
	}

	maxClientID, err := types.ParseClientID(tokens[6])
	if err != nil {
		return nil, fmt.Errorf("SERVER: couldn't parse server numeric and maximum connection count: %w", err)
	}

	description := lastParameter(tokens[7:])

	return &Server{
		Name:           tokens[1],
		HopCount:       hopCount,
		StartTimestamp: startTimestamp,
		LinkTimestamp:  linkTimestamp,
		Protocol:       protocol,
		Numeric:        maxClientID.Server,
		MaxConnections: maxClientID.Client,
		Description:    description,
	}, nil
}

func (s *Server) String() string {
	return fmt.Sprintf(
		"SERVER %s %d %d %d %s %s%s :%s",
		s.Name,
		s.HopCount,
		s.StartTimestamp,
		s.LinkTimestamp,
		s.Protocol,
		s.Numeric,
		s.MaxConnections,
		s.Description,
	)
}

type Protocol string

const (
	J10 Protocol = "J10"
	P10 Protocol = "P10"
)

func ParseProtocol(s string) (Protocol, error) {
	switch s {
	case "J10":
		return J10, nil
	case "P10":
		return P10, nil
	}

	return "", fmt.Errorf("unknown protocol: %s", s)
}

type Nick struct {
	ServerNumeric    types.ServerNumeric
	Nick             string
	HopCount         int
	CreatedTimestamp int64
	MaskUser         string
	MaskHost         string
	UserModes        string
	Account          string
	IP               string
	ClientID         types.ClientID
	Info             string
}

func ParseNick(tokens []string) (*Nick, error) {
	if len(tokens) < 10 {
		return nil, fmt.Errorf("NICK: expected at least 10 tokens; received %d", len(tokens))
	}

	serverNumeric, err := types.ParseServerNumeric(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("NICK: couldn't parse server numeric: %w", err)
	}

	if tokens[1] != "N" {
		return nil, fmt.Errorf("NICK: expected N; received %s", tokens[1])
	}

	nick := tokens[2]

	hopCount, err := strconv.Atoi(tokens[3])
	if err != nil {
		return nil, fmt.Errorf("NICK: couldn't parse hop count: %w", err)
	}

	createdTimestamp, err := strconv.ParseInt(tokens[4], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("NICK: couldn't parse created timestamp: %w", err)
	}

	maskUser := tokens[5]
	maskHost := tokens[6]

	ipIndex := 7

	userModes := ""
	account := ""

	if tokens[ipIndex][0] == '+' {
		userModes = tokens[ipIndex]
		ipIndex++

		if strings.ContainsRune(userModes, 'r') {
			account = tokens[ipIndex]
			ipIndex++
		}
	}

	ip := tokens[ipIndex]
	clientID, err := types.ParseClientID(tokens[ipIndex+1])
	if err != nil {
		return nil, fmt.Errorf("NICK: couldn't parse client ID: %w", err)
	}

	info := lastParameter(tokens[ipIndex+2:])

	return &Nick{
		ServerNumeric:    serverNumeric,
		Nick:             nick,
		HopCount:         hopCount,
		CreatedTimestamp: createdTimestamp,
		MaskUser:         maskUser,
		MaskHost:         maskHost,
		UserModes:        userModes,
		Account:          account,
		IP:               ip,
		ClientID:         clientID,
		Info:             info,
	}, nil
}

func (n *Nick) String() string {
	userModesAndAccount := ""
	if n.UserModes != "" {
		userModesAndAccount = fmt.Sprintf(" %s", n.UserModes)
		if n.Account != "" {
			userModesAndAccount += fmt.Sprintf(" %s", n.Account)
		}
	}

	return fmt.Sprintf(
		"%s N %s %d %d %s %s%s %s %s :%s",
		n.ServerNumeric,
		n.Nick,
		n.HopCount,
		n.CreatedTimestamp,
		n.MaskUser,
		n.MaskHost,
		userModesAndAccount,
		n.IP,
		n.ClientID,
		n.Info,
	)
}

type Burst struct {
	ServerNumeric    types.ServerNumeric
	Channel          string
	CreatedTimestamp int64
	ChannelModes     *types.ChannelModes
	ChannelLimit     int
	ChannelKey       string
	Members          []*types.ChannelMember
	Bans             []string
}

func ParseBurst(tokens []string) (*Burst, error) {
	if len(tokens) < 6 {
		return nil, fmt.Errorf("BURST: expected at least 7 tokens; received %d", len(tokens))
	}

	serverNumeric, err := types.ParseServerNumeric(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("BURST: couldn't parse server numeric: %w", err)
	}

	if tokens[1] != "B" {
		return nil, fmt.Errorf("BURST: expected B; received %s", tokens[1])
	}

	channel := tokens[2]

	createdTimestamp, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("BURST: couldn't parse created timestamp: %w", err)
	}

	if tokens[4][0] != '+' {
		return nil, fmt.Errorf("BURST: expected +; received %c", tokens[4][0])
	}

	channelModes, err := types.ParseChannelModes(tokens[4][1:])
	if err != nil {
		return nil, fmt.Errorf("BURST: couldn't parse channel modes: %w", err)
	}

	membersIndex := 5
	channelLimit := 0
	channelKey := ""

	if channelModes.Limit {
		channelLimit, err = strconv.Atoi(tokens[membersIndex])
		if err != nil {
			return nil, fmt.Errorf("BURST: couldn't parse channel limit: %w", err)
		}

		membersIndex++
	}

	if channelModes.Keyed {
		channelKey = tokens[membersIndex]
		membersIndex++
	}

	members, err := types.ParseChannelMembers(tokens[membersIndex])
	if err != nil {
		return nil, fmt.Errorf("BURST: couldn't parse channel members: %w", err)
	}

	var bans []string
	if len(tokens) > membersIndex+1 {
		bansString := lastParameter(tokens[membersIndex+1:])
		if len(bansString) > 0 && bansString[0] == '%' {
			bans = strings.Split(bansString[1:], " ")
		}
	}

	return &Burst{
		ServerNumeric:    serverNumeric,
		Channel:          channel,
		CreatedTimestamp: createdTimestamp,
		ChannelModes:     channelModes,
		ChannelLimit:     channelLimit,
		ChannelKey:       channelKey,
		Members:          members,
		Bans:             bans,
	}, nil
}

func (b *Burst) String() string {
	channelLimit := ""
	if b.ChannelModes.Limit {
		channelLimit = fmt.Sprintf(" %d", b.ChannelLimit)
	}

	channelKey := ""
	if b.ChannelModes.Keyed {
		channelKey = fmt.Sprintf(" %s", b.ChannelKey)
	}

	members := ""
	for i, m := range b.Members {
		if i > 0 {
			members += ","
		}
		members += m.String()
	}

	bans := ""
	if len(b.Bans) > 0 {
		bans = fmt.Sprintf(" :%%%s", strings.Join(b.Bans, " "))
	}

	return fmt.Sprintf(
		"%s B %s %d +%s%s%s %s%s",
		b.ServerNumeric,
		b.Channel,
		b.CreatedTimestamp,
		b.ChannelModes,
		channelLimit,
		channelKey,
		members,
		bans,
	)
}

type EndOfBurst struct {
	ServerNumeric types.ServerNumeric
}

func ParseEndOfBurst(tokens []string) (*EndOfBurst, error) {
	if len(tokens) != 2 {
		return nil, fmt.Errorf("END_OF_BURST: expected 2 tokens; received %d", len(tokens))
	}

	serverNumeric, err := types.ParseServerNumeric(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("END_OF_BURST: couldn't parse server numeric: %w", err)
	}

	if tokens[1] != "EB" {
		return nil, fmt.Errorf("END_OF_BURST: expected EB; received %s", tokens[1])
	}

	return &EndOfBurst{ServerNumeric: serverNumeric}, nil
}

func (e *EndOfBurst) String() string {
	return fmt.Sprintf("%s EB", e.ServerNumeric)
}

type EndOfBurstAcknowledgement struct {
	ServerNumeric types.ServerNumeric
}

func ParseEndOfBurstAcknowledgement(tokens []string) (*EndOfBurstAcknowledgement, error) {
	if len(tokens) != 2 {
		return nil, fmt.Errorf("END_OF_BURST_ACK: expected 2 tokens; received %d", len(tokens))
	}

	serverNumeric, err := types.ParseServerNumeric(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("END_OF_BURST_ACK: couldn't parse server numeric: %w", err)
	}

	if tokens[1] != "EA" {
		return nil, fmt.Errorf("END_OF_BURST_ACK: expected EA; received %s", tokens[1])
	}

	return &EndOfBurstAcknowledgement{ServerNumeric: serverNumeric}, nil
}

func (e *EndOfBurstAcknowledgement) String() string {
	return fmt.Sprintf("%s EA", e.ServerNumeric)
}

type Ping struct {
	Source types.ServerNumeric
}

func ParsePing(tokens []string) (*Ping, error) {
	if len(tokens) < 2 {
		return nil, fmt.Errorf("PING: expected at least 2 tokens; received %d", len(tokens))
	}

	source, err := types.ParseServerNumeric(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("PING: couldn't parse source: %w", err)
	}

	if tokens[1] != "G" {
		return nil, fmt.Errorf("PING: expected G; received %s", tokens[1])
	}

	return &Ping{Source: source}, nil
}

func (p *Ping) String() string {
	return fmt.Sprintf("%s G", p.Source)
}

type Pong struct {
	Source types.ServerNumeric
	Target types.ServerNumeric
}

func ParsePong(tokens []string) (*Pong, error) {
	if len(tokens) != 3 {
		return nil, fmt.Errorf("PONG: expected 3 tokens; received %d", len(tokens))
	}

	source, err := types.ParseServerNumeric(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("PONG: couldn't parse source: %w", err)
	}

	if tokens[1] != "Z" {
		return nil, fmt.Errorf("PONG: expected Z; received %s", tokens[1])
	}

	target, err := types.ParseServerNumeric(tokens[2])
	if err != nil {
		return nil, fmt.Errorf("PONG: couldn't parse target: %w", err)
	}

	return &Pong{Source: source, Target: target}, nil
}

func (p *Pong) String() string {
	return fmt.Sprintf("%s Z %s", p.Source, p.Target)
}

type Unknown struct {
	Tokens []string
}

func (u *Unknown) String() string {
	return fmt.Sprintf("UNKNOWN %s", strings.Join(u.Tokens, " "))
}

func lastParameter(ss []string) string {
	s := strings.Join(ss, " ")
	if strings.HasPrefix(s, ":") {
		return s[1:]
	}

	return s
}
