package heaps

func findKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i > -1; i-- {
		heapify(nums, 0, len(nums)-1, i)
	}

	var res int
	lastIdx := len(nums) - 1

	for i := 0; i < k; i++ {
		res = nums[0]
		nums[0], nums[lastIdx] = nums[lastIdx], nums[0]
		heapify(nums, 0, lastIdx-1, 0)
		lastIdx--
	}

	return res
}

func heapify(nums []int, low int, high int, parent int) {
	left := 2*parent + 1
	right := 2*parent + 2

	larger := parent

	if low <= left && left <= high {
		if nums[larger] < nums[left] {
			larger = left
		}
	}

	if low <= right && right <= high {
		if nums[larger] < nums[right] {
			larger = right
		}
	}

	if larger != parent {
		nums[larger], nums[parent] = nums[parent], nums[larger]
		heapify(nums, low, high, larger)
	}
}
