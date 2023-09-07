package challenges

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{
			input:  0,
			output: 0,
		},
		{
			input:  1,
			output: 1,
		},
		{
			input:  10,
			output: 55,
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Fibonacci(%d)", tt.input)
		t.Run(testName, func(t *testing.T) {
			if got := fibonacci(tt.input); got != tt.output {
				t.Errorf("For input %d, want %d, got %d", tt.input, tt.output, got)
			}
		})
	}
}
