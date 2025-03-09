package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	"pgregory.net/rapid"
)

var GeneratedPass = rapid.Custom(func(t *rapid.T) *messages.Pass {
	return &messages.Pass{
		Password: rapid.StringMatching(`^[A-Za-z][A-Za-z0-9-]{1,31}$`).Draw(t, "Password"),
	}
})
