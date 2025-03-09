package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedNick = rapid.Custom(func(t *rapid.T) *messages.Nick {
	userModes := typeGenerators.GeneratedUserModes.Draw(t, "UserModes")
	account := ""
	if userModes.Account {
		account = rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,15}$`).Draw(t, "Account")
	}

	return &messages.Nick{
		ServerNumeric:    typeGenerators.GeneratedServerNumeric.Draw(t, "ServerNumeric"),
		Nick:             rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,19}$`).Draw(t, "Nick"),
		HopCount:         rapid.IntRange(1, 10).Draw(t, "HopCount"),
		CreatedTimestamp: rapid.Int64().Draw(t, "CreatedTimestamp"),
		MaskUser:         rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,15}$`).Draw(t, "MaskUser"),
		MaskHost:         rapid.StringMatching(`^[A-Za-z][A-Za-z0-9]{1,15}$`).Draw(t, "MaskHost"),
		UserModes:        userModes,
		Account:          account,
		IP:               rapid.StringMatching(`^[A-Za-z0-9\[\]]{3,5}$`).Draw(t, "IP"),
		ClientID:         typeGenerators.GeneratedClientID.Draw(t, "ClientID"),
		Info:             rapid.StringMatching(`^[A-Za-z0-9- ]{0,64}$`).Draw(t, "Info"),
	}
})
