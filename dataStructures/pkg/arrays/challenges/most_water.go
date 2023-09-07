package challenges

func maxArea(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}

	var area int
	a := 0
	b := n - 1

	for a < b {
		var length int
		if height[a] < height[b] {
			length = height[a]
		} else {
			length = height[b]
		}
		width := b - a
		tempArea := calculateArea(length, width)
		if tempArea > area {
			area = tempArea
		}

		if height[a] < height[b] {
			a++
		} else {
			b--
		}
	}

	return area
}

func calculateArea(length, width int) int {
	return length * width
}
