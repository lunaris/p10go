package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedPrivmsg = rapid.Custom(func(t *rapid.T) *messages.Privmsg {
	return &messages.Privmsg{
		Source:  typeGenerators.GeneratedClientID.Draw(t, "Source"),
		Target:  typeGenerators.GeneratedClientID.Draw(t, "Target"),
		Message: rapid.StringMatching(`^[A-Za-z0-9- ]{0,64}$`).Draw(t, "Message"),
	}
})
