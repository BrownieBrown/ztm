package challenges

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{7, 1, 5, 3, 6, 4},
			output: 5,
		},
		{
			input:  []int{7, 6, 4, 3, 1},
			output: 0,
		},
	}
	for _, test := range tests {
		testName := fmt.Sprintf("MaxProfit(%d)", test.input)
		t.Run(testName, func(t *testing.T) {
			if got := maxProfit(test.input); got != test.output {
				t.Errorf("For input %d, want %d, got %d", test.input, test.output, got)
			}
		})
	}
}
