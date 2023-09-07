package challenges

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{
			input:  2,
			output: 2,
		},
		{
			input:  3,
			output: 3,
		},
	}
	for _, test := range tests {
		testName := fmt.Sprintf("ClimbStairs(%d)", test.input)
		t.Run(testName, func(t *testing.T) {
			if got := climbStairs(test.input); got != test.output {
				t.Errorf("For input %d, want %d, got %d", test.input, test.output, got)
			}
		})
	}
}
