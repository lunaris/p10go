package test

import (
	"strings"
	"testing"

	"github.com/lunaris/p10go/pkg/messages"
	"github.com/lunaris/p10go/test/generators"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestPassMessagesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedPass.Draw(t, "ServerNumeric")

		actual, err := messages.ParsePass(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestServerMessagesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedServer.Draw(t, "Server")

		actual, err := messages.ParseServer(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestNickMessagesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedNick.Draw(t, "Nick")

		actual, err := messages.ParseNick(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestBurstMessagesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedBurst.Draw(t, "Burst")

		actual, err := messages.ParseBurst(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestPingMessagesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedPing.Draw(t, "Ping")

		actual, err := messages.ParsePing(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestPongMessagesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedPong.Draw(t, "Pong")

		actual, err := messages.ParsePong(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestJoinMessagesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedJoin.Draw(t, "Join")

		actual, err := messages.ParseJoin(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestChannelModeMessagesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedChannelMode.Draw(t, "ChannelMode")

		t.Logf("expected: %s", expected.String())

		actual, err := messages.ParseChannelMode(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestUserModeMessagesRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expected := generators.GeneratedUserMode.Draw(t, "UserMode")

		t.Logf("expected: %s", expected.String())

		actual, err := messages.ParseUserMode(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
