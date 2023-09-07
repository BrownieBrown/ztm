package challenges

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "abba",
			output: true,
		},
		{
			input:  "abb",
			output: false,
		},
		{
			input:  "race a car",
			output: false,
		},
		{
			input:  "A man, a plan, a canal: Panama",
			output: true,
		},
		{
			input:  "aabaa",
			output: true,
		},
		{
			input:  "",
			output: true,
		},
		{
			input:  " ",
			output: true,
		},
		{
			input:  "ab",
			output: false,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("isPalindrome(%s)", test.input)
		t.Run(testName, func(t *testing.T) {
			if got, want := isPalindrome(test.input), test.output; got != want {
				t.Errorf("For input %s, want %v, got %v", test.input, test.output, got)
			}
		})
	}
}
