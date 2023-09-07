package quick

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := partition(arr)
	quickSort(arr[:pivotIndex])   // Sort the left part
	quickSort(arr[pivotIndex+1:]) // Sort the right part

	return arr
}

func partition(arr []int) int {
	pivot := arr[len(arr)/2] // Select the pivot (middle element in this case)
	i := 0
	j := len(arr) - 1

	for {
		// Find the first element greater than the pivot
		for arr[i] < pivot {
			i++
		}

		// Find the first element less than the pivot
		for arr[j] > pivot {
			j--
		}

		// If pointers have crossed, the partition is done
		if i >= j {
			return j
		}

		// Swap elements to place them on the correct side of the pivot
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
