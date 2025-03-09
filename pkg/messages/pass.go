package messages

import "fmt"

type Pass struct {
	Password string
}

func ParsePass(tokens []string) (*Pass, error) {
	if len(tokens) != 2 {
		return nil, fmt.Errorf("PASS: expected 2 tokens; received %d", len(tokens))
	}

	if tokens[0] != "PASS" {
		return nil, fmt.Errorf("PASS: expected PASS; received %s", tokens[0])
	}

	return &Pass{Password: lastParameter(tokens[1:])}, nil
}

func (p *Pass) String() string {
	return fmt.Sprintf("PASS :%s", p.Password)
}
