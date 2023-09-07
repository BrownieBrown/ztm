package selection

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIdx := i // Keep track of the index of the smallest element
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i] // Swap once per iteration of the outer loop
	}
	return arr
}
