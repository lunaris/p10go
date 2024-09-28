package test

import (
	"testing"

	"github.com/lunaris/p10go/pkg/types"
	"github.com/lunaris/p10go/test/generators"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestServerNumericsRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedServerNumeric.Draw(t, "ServerNumeric")

		actual, err := types.ParseServerNumeric(string(expected))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestClientNumericsRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedClientNumeric.Draw(t, "ClientNumeric")

		actual, err := types.ParseClientNumeric(string(expected))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestClientIDsRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedClientID.Draw(t, "ClientID")

		actual, err := types.ParseClientID(expected.String())

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestChannelModesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedChannelModes.Draw(t, "ChannelModes")

		actual, err := types.ParseChannelModes(expected.String())

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestUserModesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedUserModes.Draw(t, "UserModes")

		actual, err := types.ParseUserModes(expected.String())

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestChannelUserModesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedChannelUserModes.Draw(t, "ChannelUserModes")

		actual, err := types.ParseChannelUserModes(expected.String())

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestChannelMembersRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedChannelMember.Draw(t, "ChannelMember")

		actual, err := types.ParseChannelMember(expected.String())

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
