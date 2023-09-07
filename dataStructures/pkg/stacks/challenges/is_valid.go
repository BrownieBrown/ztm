package challenges

func isValid(s string) bool {
	if s == "" {
		return true
	}

	var stack []rune
	hashMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, current := range s {
		// If current is a closing brace/bracket/parenthesis
		if expected, ok := hashMap[current]; ok {
			// If the stack is empty or doesn't match the expected opening brace
			if len(stack) == 0 || stack[len(stack)-1] != expected {
				return false
			}
			// Pop from stack
			stack = stack[:len(stack)-1]
		} else {
			// Else push to stack
			stack = append(stack, current)
		}
	}

	// If stack is empty, all braces/brackets/parentheses were matched
	return len(stack) == 0
}
