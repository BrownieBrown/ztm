package merge

func mergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	midIndex := n / 2
	leftArr := mergeSort(arr[:midIndex])  // Recursively call mergeSort on left half
	rightArr := mergeSort(arr[midIndex:]) // Recursively call mergeSort on right half

	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) []int {
	var result []int
	i, j := 0, 0

	// Compare the elements in the two arrays and add the smaller one to the result
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] < rightArr[j] {
			result = append(result, leftArr[i])
			i++
		} else {
			result = append(result, rightArr[j])
			j++
		}
	}

	// If there are any remaining elements in the left or right array, append them to the result
	for ; i < len(leftArr); i++ {
		result = append(result, leftArr[i])
	}
	for ; j < len(rightArr); j++ {
		result = append(result, rightArr[j])
	}

	return result
}
