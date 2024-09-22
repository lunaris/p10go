package messages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lunaris/p10go/numerics"
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
	case "N", "NICK":
		return ParseNick(tokens)
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
	Numeric        numerics.ServerNumeric
	MaxConnections numerics.ClientNumeric
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

	maxClientID, err := numerics.ParseClientID(tokens[6])
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
	ServerNumeric    numerics.ServerNumeric
	Nick             string
	HopCount         int
	CreatedTimestamp int64
	MaskUser         string
	MaskHost         string
	UserModes        string
	Account          string
	IP               string
	ClientID         numerics.ClientID
	Info             string
}

func ParseNick(tokens []string) (*Nick, error) {
	if len(tokens) < 10 {
		return nil, fmt.Errorf("NICK: expected at least 10 tokens; received %d", len(tokens))
	}

	serverNumeric, err := numerics.ParseServerNumeric(tokens[0])
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
	clientID, err := numerics.ParseClientID(tokens[ipIndex+1])
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
