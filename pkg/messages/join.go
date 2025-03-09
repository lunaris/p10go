package messages

import (
	"fmt"
	"strconv"

	"github.com/lunaris/p10go/pkg/types"
)

type Join struct {
	ClientID  types.ClientID
	Channel   string
	Timestamp int64
}

func ParseJoin(tokens []string) (*Join, error) {
	clientID, err := types.ParseClientID(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("JOIN: couldn't parse client ID: %w", err)
	}

	if tokens[1] != "J" {
		return nil, fmt.Errorf("JOIN: expected J; received %s", tokens[1])
	}

	channel := tokens[2]

	timestamp, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("JOIN: couldn't parse timestamp: %w", err)
	}

	return &Join{ClientID: clientID, Channel: channel, Timestamp: timestamp}, nil
}

func (j *Join) String() string {
	return fmt.Sprintf("%s J %s %d", j.ClientID, j.Channel, j.Timestamp)
}
