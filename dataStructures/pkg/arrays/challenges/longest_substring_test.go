package challenges

import (
	"fmt"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "dvdf",
			output: 3,
		},
		{
			input:  "abcabcbb",
			output: 3,
		},
		{
			input:  "bbbbb",
			output: 1,
		},
		{
			input:  "pwwkew",
			output: 3,
		},
		{
			input:  "au",
			output: 2,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("lengthOfLongestSubstring(%s)", test.input)
		t.Run(testName, func(t *testing.T) {
			if got, want := lengthOfLongestSubstring(test.input), test.output; got != want {
				t.Errorf("For input %s, want %v, got %v", test.input, test.output, got)
			}
		})
	}
}
