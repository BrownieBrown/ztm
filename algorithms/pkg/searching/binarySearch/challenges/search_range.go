package challenges

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1} // default result if target is not found

	// Find the starting index of target
	leftIndex := search(nums, target, true)
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return result
	}

	result[0] = leftIndex
	result[1] = search(nums, target, false) - 1

	return result
}

func search(nums []int, target int, findFirst bool) int {
	low, high := 0, len(nums)
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > target || (findFirst && nums[mid] == target) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
