package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedEndOfBurst = rapid.Custom(func(t *rapid.T) *messages.EndOfBurst {
	return &messages.EndOfBurst{
		ServerNumeric: typeGenerators.GeneratedServerNumeric.Draw(t, "ServerNumeric"),
	}
})
