package repl

import "testing"

func TestCleanInpput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "    hello ",
			expected: []string{"hello"},
		},
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hello       world     ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match")
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]

			if actualWord != expectedWord {
				t.Errorf("word doesn't match")
			}
		}
	}
}
