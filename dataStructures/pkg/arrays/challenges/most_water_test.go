package challenges

import (
	"fmt"
	"testing"
)

func TestMostArea(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
		{
			input:  []int{2, 3, 4, 5, 18, 17, 6},
			output: 17,
		},
		{
			input:  []int{1, 1},
			output: 1,
		},
		{
			input:  []int{},
			output: 0,
		},
		{
			input:  []int{1},
			output: 0,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("maxArea(%d)", test.input)
		t.Run(testName, func(t *testing.T) {
			if got, want := maxArea(test.input), test.output; got != want {
				t.Errorf("For input %d, want %d, got %d", test.input, test.output, got)
			}
		})
	}
}
