package numerics

import (
	"fmt"
	"regexp"
)

type ServerNumeric string

var serverNumericPattern = regexp.MustCompile(`^[A-Za-z0-9\[\]]{2}$`)

func ParseServerNumeric(s string) (ServerNumeric, error) {
	if !serverNumericPattern.MatchString(s) {
		return "", fmt.Errorf("invalid server numeric: %s", s)
	}

	return ServerNumeric(s), nil
}

type ClientNumeric string

var clientNumericPattern = regexp.MustCompile(`^[A-Za-z0-9\[\]]{3}$`)

func ParseClientNumeric(s string) (ClientNumeric, error) {
	if !clientNumericPattern.MatchString(s) {
		return "", fmt.Errorf("invalid client numeric: %s", s)
	}

	return ClientNumeric(s), nil
}

type ClientID struct {
	Server ServerNumeric
	Client ClientNumeric
}

func ParseClientID(s string) (ClientID, error) {
	if len(s) != 5 {
		return ClientID{}, fmt.Errorf("invalid client ID: %s", s)
	}

	server, err := ParseServerNumeric(s[:2])
	if err != nil {
		return ClientID{}, fmt.Errorf("couldn't parse server numeric: %w", err)
	}

	client, err := ParseClientNumeric(s[2:])
	if err != nil {
		return ClientID{}, fmt.Errorf("couldn't parse client numeric: %w", err)
	}

	return ClientID{
		Server: server,
		Client: client,
	}, nil
}
