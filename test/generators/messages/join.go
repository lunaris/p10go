package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedChannelName = rapid.StringMatching(`^[#][A-Za-z][A-Za-z0-9-]{1,31}$`)

var GeneratedJoin = rapid.Custom(func(t *rapid.T) *messages.Join {
	return &messages.Join{
		ClientID:  typeGenerators.GeneratedClientID.Draw(t, "ClientID"),
		Channel:   GeneratedChannelName.Draw(t, "Channel"),
		Timestamp: rapid.Int64().Draw(t, "Timestamp"),
	}
})
