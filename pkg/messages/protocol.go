package messages

import "fmt"

type Protocol string

const (
	J10 Protocol = "J10"
	P10 Protocol = "P10"
)

func ParseProtocol(s string) (Protocol, error) {
	switch s {
	case "J10":
		return J10, nil
	case "P10":
		return P10, nil
	}

	return "", fmt.Errorf("unknown protocol: %s", s)
}
