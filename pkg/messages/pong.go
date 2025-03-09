package messages

import (
	"fmt"

	"github.com/lunaris/p10go/pkg/types"
)

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
