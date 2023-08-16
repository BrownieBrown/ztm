package arrays

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		compliment := target - num
		if index, found := m[compliment]; found {
			return []int{index, i}
		} else {
			m[num] = i
		}
	}
	return nil
}
