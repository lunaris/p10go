package messages

import (
	"fmt"

	"github.com/lunaris/p10go/pkg/types"
)

type Privmsg struct {
	Source  types.ClientID
	Target  types.ClientID
	Message string
}

func ParsePrivmsg(tokens []string) (*Privmsg, error) {
	if len(tokens) < 4 {
		return nil, fmt.Errorf("PRIVMSG: expected at least 4 tokens; received %d", len(tokens))
	}

	source, err := types.ParseClientID(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("PRIVMSG: couldn't parse source: %w", err)
	}

	if tokens[1] != "P" {
		return nil, fmt.Errorf("PRIVMSG: expected P; received %s", tokens[1])
	}

	target, err := types.ParseClientID(tokens[2])
	if err != nil {
		return nil, fmt.Errorf("PRIVMSG: couldn't parse target: %w", err)
	}

	message := lastParameter(tokens[3:])

	return &Privmsg{
		Source:  source,
		Target:  target,
		Message: message,
	}, nil
}

func (p *Privmsg) String() string {
	return fmt.Sprintf("%s P %s :%s", p.Source, p.Target, p.Message)
}
