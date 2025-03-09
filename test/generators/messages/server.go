package messages

import (
	"github.com/lunaris/p10go/pkg/messages"
	typeGenerators "github.com/lunaris/p10go/test/generators/types"
	"pgregory.net/rapid"
)

var GeneratedServer = rapid.Custom(func(t *rapid.T) *messages.Server {
	return &messages.Server{
		Name:           rapid.StringMatching(`^[A-Za-z][A-Za-z0-9-]{1,31}$`).Draw(t, "Name"),
		HopCount:       rapid.IntRange(1, 10).Draw(t, "HopCount"),
		StartTimestamp: rapid.Int64().Draw(t, "StartTimestamp"),
		LinkTimestamp:  rapid.Int64().Draw(t, "LinkTimestamp"),
		Protocol: rapid.SampledFrom([]messages.Protocol{
			messages.J10,
			messages.J10,
		}).Draw(t, "Protocol"),
		Numeric:        typeGenerators.GeneratedServerNumeric.Draw(t, "Numeric"),
		MaxConnections: typeGenerators.GeneratedClientNumeric.Draw(t, "MaxConnections"),
		Description:    rapid.StringMatching(`^[A-Za-z0-9- ]{0,64}$`).Draw(t, "Description"),
	}
})
