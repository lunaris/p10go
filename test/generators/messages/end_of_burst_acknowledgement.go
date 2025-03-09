package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedEndOfBurstAcknowledgement = rapid.Custom(func(t *rapid.T) *messages.EndOfBurstAcknowledgement {
	return &messages.EndOfBurstAcknowledgement{
		ServerNumeric: typeGenerators.GeneratedServerNumeric.Draw(t, "ServerNumeric"),
	}
})
