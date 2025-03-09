package messages

import (
	"fmt"
	"strconv"

	"github.com/lunaris/p10go/pkg/types"
)

type Nick struct {
	ServerNumeric    types.ServerNumeric
	Nick             string
	HopCount         int
	CreatedTimestamp int64
	MaskUser         string
	MaskHost         string
	UserModes        types.UserModes
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

	var userModes types.UserModes
	account := ""

	if tokens[ipIndex][0] == '+' {
		userModes, err = types.ParseUserModes(tokens[ipIndex][1:])
		if err != nil {
			return nil, fmt.Errorf("NICK: couldn't parse user modes: %w", err)
		}

		ipIndex++

		if userModes.Account {
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
	userModes := n.UserModes.String()
	if len(userModes) > 0 {
		userModesAndAccount = " +" + userModes
		if n.Account != "" {
			userModesAndAccount += " " + n.Account
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
