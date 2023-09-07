package challenges

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	arr := []rune(s)
	lastOccurrence := make(map[rune]int) // HashMap to keep track of the last index at which each character appeared
	longestSubstring := 0
	startIndex := 0 // Start index of the current substring

	for endIndex, char := range arr {
		// If char is already in the map, then move the start index of the current substring
		if lastIdx, exists := lastOccurrence[char]; exists && lastIdx >= startIndex {
			startIndex = lastIdx + 1
		}

		// Update the result if we get a larger window
		longestSubstring = Max(longestSubstring, endIndex-startIndex+1)

		// Update last occurrence of the character
		lastOccurrence[char] = endIndex
	}

	return longestSubstring
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
