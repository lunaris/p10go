package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedPong = rapid.Custom(func(t *rapid.T) *messages.Pong {
	return &messages.Pong{
		Source: typeGenerators.GeneratedServerNumeric.Draw(t, "Source"),
		Target: typeGenerators.GeneratedServerNumeric.Draw(t, "Target"),
	}
})
