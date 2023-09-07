package challenges

import (
	"fmt"
	"testing"
)

func TestTrapRainwater(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			output: 6,
		},
		{
			input:  []int{4, 2, 0, 3, 2, 5},
			output: 9,
		},
		{
			input:  []int{1, 1},
			output: 0,
		},
		{
			input:  []int{},
			output: 0,
		},
		{
			input:  []int{1, 0},
			output: 0,
		},
		{
			input:  []int{0, 1},
			output: 0,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("trapRain(%d)", test.input)
		t.Run(testName, func(t *testing.T) {
			if got, want := trap(test.input), test.output; got != want {
				t.Errorf("For input %d, want %d, got %d", test.input, test.output, got)
			}
		})
	}
}
