package test

import (
	"strings"
	"testing"

	"github.com/lunaris/p10go/pkg/messages"
	messageGenerators "github.com/lunaris/p10go/test/generators/messages"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestPassMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedPass.Draw(t, "ServerNumeric")

		actual, err := messages.ParsePass(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestServerMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedServer.Draw(t, "Server")

		actual, err := messages.ParseServer(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestNickMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedNick.Draw(t, "Nick")

		actual, err := messages.ParseNick(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestBurstMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedBurst.Draw(t, "Burst")

		actual, err := messages.ParseBurst(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestPingMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedPing.Draw(t, "Ping")

		actual, err := messages.ParsePing(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestPongMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedPong.Draw(t, "Pong")

		actual, err := messages.ParsePong(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestJoinMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedJoin.Draw(t, "Join")

		actual, err := messages.ParseJoin(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestChannelModeMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedChannelMode.Draw(t, "ChannelMode")

		t.Logf("expected: %s", expected.String())

		actual, err := messages.ParseChannelMode(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestUserModeMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedUserMode.Draw(t, "UserMode")

		t.Logf("expected: %s", expected.String())

		actual, err := messages.ParseUserMode(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestPrivmsgMessagesRoundtrip(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(t *rapid.T) {
		expected := messageGenerators.GeneratedPrivmsg.Draw(t, "Privmsg")

		t.Logf("expected: %s", expected.String())

		actual, err := messages.ParsePrivmsg(strings.Split(expected.String(), " "))

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
