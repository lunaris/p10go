package messages

import (
	"fmt"

	"github.com/lunaris/p10go/pkg/types"
)

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
