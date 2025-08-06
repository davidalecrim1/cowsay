package cowsay

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/expected_cowsay_001.txt
var expectedSingleLineCowSay string

//go:embed testdata/expected_cowsay_002.txt
var expectedHelloWorldCowSay string

//go:embed testdata/expected_cowsay_003.txt
var expectedLineBreakCowSay string

//go:embed testdata/expected_cowsay_004.txt
var expectedTabCowSay string

func TestGenerateCowSay(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "single word",
			input:    []string{"hello"},
			expected: expectedSingleLineCowSay,
		},
		{
			name:     "simple text",
			input:    []string{"hello, world!"},
			expected: expectedHelloWorldCowSay,
		},
		{
			name:     "line break",
			input:    []string{"david", "alecrim"},
			expected: expectedLineBreakCowSay,
		},
		{
			name: "handle text with tabs",
			input: []string{
				"When I was in school, I cheated on my metaphysics exam: I looked into",
				"the soul of the boy sitting next to me.",
				"\t\t-- Woody Allen",
			},
			expected: expectedTabCowSay,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := GenerateCowSay(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
