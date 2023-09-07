package challenges

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{
			input:  5,
			output: 120,
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Factorial(%d)", tt.input)
		t.Run(testName, func(t *testing.T) {
			if got := findFactorial(tt.input); got != tt.output {
				t.Errorf("For input %d, want %d, got %d", tt.input, tt.output, got)
			}
		})
	}
}
