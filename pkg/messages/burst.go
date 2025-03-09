package messages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lunaris/p10go/pkg/types"
)

type Burst struct {
	ServerNumeric    types.ServerNumeric
	Channel          string
	CreatedTimestamp int64
	ChannelModes     types.ChannelModes
	ChannelLimit     int
	ChannelKey       string
	Members          []types.ChannelMember
	Bans             []string
}

func ParseBurst(tokens []string) (*Burst, error) {
	if len(tokens) < 5 {
		return nil, fmt.Errorf("BURST: expected at least 5 tokens; received %d", len(tokens))
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

	membersIndex := 4
	var channelModes types.ChannelModes
	channelLimit := 0
	channelKey := ""

	if tokens[4][0] == '+' {
		channelModes, _, err = types.ParseChannelModes(tokens[4][1:])
		if err != nil {
			return nil, fmt.Errorf("BURST: couldn't parse channel modes: %w", err)
		}

		membersIndex++

		// In a burst, limit always comes before key, so we don't need to use the
		// parameterized modes return value of ParseChannelModes.
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
	}

	members, err := types.ParseChannelMembers(tokens[membersIndex])
	if err != nil {
		return nil, fmt.Errorf("BURST: couldn't parse channel members: %w", err)
	}

	bans := []string{}
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
	channelModes := b.ChannelModes.String()
	channelLimit := ""
	channelKey := ""
	if len(channelModes) > 0 {
		channelModes = " +" + channelModes
		if b.ChannelModes.Limit {
			channelLimit = fmt.Sprintf(" %d", b.ChannelLimit)
		}

		if b.ChannelModes.Keyed {
			channelKey = fmt.Sprintf(" %s", b.ChannelKey)
		}
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
		"%s B %s %d%s%s%s %s%s",
		b.ServerNumeric,
		b.Channel,
		b.CreatedTimestamp,
		channelModes,
		channelLimit,
		channelKey,
		members,
		bans,
	)
}
