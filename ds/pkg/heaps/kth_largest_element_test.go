package heaps

import (
	"testing"
)

type Input struct {
	nums []int
	k    int
}

type Output struct {
	result int
}

func TestKthLargestElement(t *testing.T) {
	tests := []struct {
		input  Input
		output Output
	}{
		{
			input: Input{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			output: Output{
				result: 5,
			},
		},
		{
			input: Input{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			output: Output{
				result: 4,
			},
		},
	}

	for _, tt := range tests {
		got := findKthLargest(tt.input.nums, tt.input.k)
		if got != tt.output.result {
			t.Errorf("For nums: %v and k: %d, expected %d but got %d", tt.input.nums, tt.input.k, tt.output.result, got)
		}
	}
}
