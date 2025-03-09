package messages

import (
	"fmt"
)

type Message interface {
	String() string
}

func Parse(tokens []string) (Message, error) {
	if len(tokens) < 2 {
		return nil, fmt.Errorf("expected at least 2 tokens; received %d", len(tokens))
	}

	switch tokens[0] {
	case "PASS":
		return ParsePass(tokens)
	case "SERVER":
		return ParseServer(tokens)
	}

	switch tokens[1] {
	case "B":
		return ParseBurst(tokens)
	case "EA":
		return ParseEndOfBurstAcknowledgement(tokens)
	case "EB":
		return ParseEndOfBurst(tokens)
	case "G":
		return ParsePing(tokens)
	case "J":
		return ParseJoin(tokens)
	case "M", "OM":
		if tokens[2][0] == '#' {
			return ParseChannelMode(tokens)
		} else {
			return ParseUserMode(tokens)
		}
	case "N":
		return ParseNick(tokens)
	case "P":
		return ParsePrivmsg(tokens)
	case "Z":
		return ParsePong(tokens)
	}

	return &Unknown{Tokens: tokens}, nil
}
