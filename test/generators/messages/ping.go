package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedPing = rapid.Custom(func(t *rapid.T) *messages.Ping {
	return &messages.Ping{
		Source: typeGenerators.GeneratedServerNumeric.Draw(t, "Source"),
	}
})
