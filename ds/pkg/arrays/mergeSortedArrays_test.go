package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		arr1       []int
		arr2       []int
		mergeArray []int
	}{
		{arr1: []int{0, 1, 2, 5, 6}, arr2: []int{2, 3, 5, 7}, mergeArray: []int{0, 1, 2, 2, 3, 5, 5, 6, 7}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Merge %v and %v", tt.arr1, tt.arr2)
		t.Run(testName, func(t *testing.T) {
			got, err := mergeSortedArrays(tt.arr1, tt.arr2)
			if err != nil {
				t.Errorf("Error while merging arrays: %v", err)
			}

			if !reflect.DeepEqual(got, tt.mergeArray) {
				t.Errorf("Expected %v, got %v", tt.mergeArray, got)
			}
		})
	}
}
