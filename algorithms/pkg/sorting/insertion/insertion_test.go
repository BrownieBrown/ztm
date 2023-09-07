package insertion

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{64, 34, 25, 12, 22, 11, 90},
			output: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			input:  []int{5, 1, 4, 2, 8},
			output: []int{1, 2, 4, 5, 8},
		},
		{
			input:  []int{3, 0, 2, 5, -1, 4, 1},
			output: []int{-1, 0, 1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("InsertionSort(%v)", tt.input)
		t.Run(testName, func(t *testing.T) {
			if got := insertionSort(tt.input); !reflect.DeepEqual(got, tt.output) {
				t.Errorf("For input %v, want %v, got %v", tt.input, tt.output, got)
			}
		})
	}
}
