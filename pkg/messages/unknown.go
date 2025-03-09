package messages

import (
	"fmt"
	"strings"
)

type Unknown struct {
	Tokens []string
}

func (u *Unknown) String() string {
	return fmt.Sprintf("UNKNOWN %s", strings.Join(u.Tokens, " "))
}

func lastParameter(ss []string) string {
	s := strings.Join(ss, " ")
	if strings.HasPrefix(s, ":") {
		return s[1:]
	}

	return s
}
