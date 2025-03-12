package test

import (
	"testing"

	"github.com/lunaris/p10go/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestParseChannelMembers(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input    string
		expected []types.ChannelMember
	}{
		{
			input: "ABCDE",
			expected: []types.ChannelMember{
				{
					ClientID: types.ClientID{Server: "AB", Client: "CDE"},
				},
			},
		},
		{
			input: "ABCDE:v",
			expected: []types.ChannelMember{
				{
					ClientID: types.ClientID{Server: "AB", Client: "CDE"},
					Modes:    types.ChannelUserModes{Voice: true},
				},
			},
		},
		{
			input: "ABCDE:o",
			expected: []types.ChannelMember{
				{
					ClientID: types.ClientID{Server: "AB", Client: "CDE"},
					Modes:    types.ChannelUserModes{Op: true},
				},
			},
		},
		{
			input: "ABCDE:ov",
			expected: []types.ChannelMember{
				{
					ClientID: types.ClientID{Server: "AB", Client: "CDE"},
					Modes:    types.ChannelUserModes{Op: true, Voice: true},
				},
			},
		},
		{
			input: "ABCDE,FGHIJ",
			expected: []types.ChannelMember{
				{
					ClientID: types.ClientID{Server: "AB", Client: "CDE"},
				},
				{
					ClientID: types.ClientID{Server: "FG", Client: "HIJ"},
				},
			},
		},
		{
			input: "ABCDE:v,FGHIJ",
			expected: []types.ChannelMember{
				{
					ClientID: types.ClientID{Server: "AB", Client: "CDE"},
					Modes:    types.ChannelUserModes{Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "FG", Client: "HIJ"},
					Modes:    types.ChannelUserModes{Voice: true},
				},
			},
		},
		{
			input: "ABCDE:o,FGHIJ",
			expected: []types.ChannelMember{
				{
					ClientID: types.ClientID{Server: "AB", Client: "CDE"},
					Modes:    types.ChannelUserModes{Op: true},
				},
				{
					ClientID: types.ClientID{Server: "FG", Client: "HIJ"},
					Modes:    types.ChannelUserModes{Op: true},
				},
			},
		},
		{
			input: "ABCDE,FGHIJ,KLMNO:v,PQRST,UVWXY,ZABCD:o,EFGHI:ov,JKLMN,OPQRS",
			expected: []types.ChannelMember{
				{
					ClientID: types.ClientID{Server: "AB", Client: "CDE"},
				},
				{
					ClientID: types.ClientID{Server: "FG", Client: "HIJ"},
				},
				{
					ClientID: types.ClientID{Server: "KL", Client: "MNO"},
					Modes:    types.ChannelUserModes{Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "PQ", Client: "RST"},
					Modes:    types.ChannelUserModes{Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "UV", Client: "WXY"},
					Modes:    types.ChannelUserModes{Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "ZA", Client: "BCD"},
					Modes:    types.ChannelUserModes{Op: true},
				},
				{
					ClientID: types.ClientID{Server: "EF", Client: "GHI"},
					Modes:    types.ChannelUserModes{Op: true, Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "JK", Client: "LMN"},
					Modes:    types.ChannelUserModes{Op: true, Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "OP", Client: "QRS"},
					Modes:    types.ChannelUserModes{Op: true, Voice: true},
				},
			},
		},
		{
			input: "ABCDE,FGHIJ:v,KLMNO:o,PQRST:ov,UVWXY:v,ZABCD:o,EFGHI:ov,JKLMN,OPQRS",
			expected: []types.ChannelMember{
				{
					ClientID: types.ClientID{Server: "AB", Client: "CDE"},
				},
				{
					ClientID: types.ClientID{Server: "FG", Client: "HIJ"},
					Modes:    types.ChannelUserModes{Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "KL", Client: "MNO"},
					Modes:    types.ChannelUserModes{Op: true},
				},
				{
					ClientID: types.ClientID{Server: "PQ", Client: "RST"},
					Modes:    types.ChannelUserModes{Op: true, Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "UV", Client: "WXY"},
					Modes:    types.ChannelUserModes{Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "ZA", Client: "BCD"},
					Modes:    types.ChannelUserModes{Op: true},
				},
				{
					ClientID: types.ClientID{Server: "EF", Client: "GHI"},
					Modes:    types.ChannelUserModes{Op: true, Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "JK", Client: "LMN"},
					Modes:    types.ChannelUserModes{Op: true, Voice: true},
				},
				{
					ClientID: types.ClientID{Server: "OP", Client: "QRS"},
					Modes:    types.ChannelUserModes{Op: true, Voice: true},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.input, func(t *testing.T) {
			t.Parallel()

			actual, err := types.ParseChannelMembers(c.input)

			assert.NoError(t, err)
			assert.Equal(t, c.expected, actual)
		})
	}
}
