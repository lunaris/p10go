package messages

import (
	"fmt"

	"github.com/lunaris/p10go/pkg/types"
)

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
