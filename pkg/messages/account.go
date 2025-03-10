package messages

import (
	"fmt"
	"strconv"

	"github.com/lunaris/p10go/pkg/types"
)

type Account struct {
	Source      types.ServerNumeric
	Target      types.ClientID
	AccountName string
	Timestamp   int64
}

func ParseAccount(tokens []string) (*Account, error) {
	if len(tokens) < 5 {
		return nil, fmt.Errorf("ACCOUNT: expected 5 tokens; received %d", len(tokens))
	}

	source, err := types.ParseServerNumeric(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("ACCOUNT: couldn't parse source: %w", err)
	}

	if tokens[1] != "AC" {
		return nil, fmt.Errorf("ACCOUNT: expected AC; received %s", tokens[1])
	}

	target, err := types.ParseClientID(tokens[2])
	if err != nil {
		return nil, fmt.Errorf("ACCOUNT: couldn't parse target: %w", err)
	}

	timestamp, err := strconv.ParseInt(tokens[4], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("ACCOUNT: couldn't parse timestamp: %w", err)
	}

	return &Account{
		Source:      source,
		Target:      target,
		AccountName: tokens[3],
		Timestamp:   timestamp,
	}, nil
}

func (a *Account) String() string {
	return fmt.Sprintf("%s AC %s %s %d 0", a.Source, a.Target, a.AccountName, a.Timestamp)
}
