package challenges

func minRemoveToMakeValid(s string) string {
	if s == "" {
		return s
	}

	var openingStack, closingStack []int

	for index, char := range s {
		if isOpening(char) {
			openingStack = append(openingStack, index)
		} else if isClosing(char) {
			if len(openingStack) != 0 {
				openingStack = openingStack[:len(openingStack)-1]
			} else {
				closingStack = append(closingStack, index)
			}
		}
	}

	// Convert the leftover indices in stacks to a set for O(1) lookup
	toRemove := map[int]struct{}{}
	for _, idx := range openingStack {
		toRemove[idx] = struct{}{}
	}
	for _, idx := range closingStack {
		toRemove[idx] = struct{}{}
	}

	// Construct the final result, skipping over the indices to remove
	var result []rune
	for idx, char := range s {
		if _, exists := toRemove[idx]; !exists {
			result = append(result, char)
		}
	}

	return string(result)
}

func isOpening(c rune) bool {
	return c == '('
}

func isClosing(c rune) bool {
	return c == ')'
}
