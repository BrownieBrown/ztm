package challenges

import (
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func isPalindrome(s string) bool {
	// Clear out non-alphanumeric characters
	str := clearString(s)

	// Remove whitespaces and convert to lowercase
	str = strings.Join(strings.Fields(str), "")
	str = strings.ToLower(str)

	// Convert to rune array for Unicode handling
	arr := []rune(str)
	n := len(arr)

	// Handle trivial cases
	if n < 2 {
		return true
	}

	left, right := 0, n-1

	// Check palindrome property
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}

	return true
}
