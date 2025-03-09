package messages

import (
	"fmt"

	"github.com/lunaris/p10go/pkg/types"
)

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
