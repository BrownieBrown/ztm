package challenges

import (
	"fmt"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		output bool
	}{
		{
			input1: "ab#c",
			input2: "ad#c",
			output: true,
		},
		{
			input1: "ab##",
			input2: "c#d#",
			output: true,
		},
		{
			input1: "ab#c",
			input2: "ad#c",
			output: true,
		},
		{
			input1: "a#c",
			input2: "b",
			output: false,
		},
		{
			input1: "",
			input2: "",
			output: true,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("backspaceCompare(%s,%s)", test.input1, test.input2)
		t.Run(testName, func(t *testing.T) {
			if got, want := backspaceCompare(test.input1, test.input2), test.output; got != want {
				t.Errorf("For inputs %s and %s, want %v, got %v", test.input1, test.input2, test.output, got)
			}
		})
	}
}
