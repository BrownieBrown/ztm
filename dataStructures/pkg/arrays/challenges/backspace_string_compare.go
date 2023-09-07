package challenges

func backspaceCompare(s string, t string) bool {
	queue1 := []rune(s)
	queue2 := []rune(t)

	removeBackspace(&queue1)
	removeBackspace(&queue2)

	return compareQueues(queue1, queue2)
}

func compareQueues(queue1, queue2 []rune) bool {
	if len(queue1) != len(queue2) {
		return false
	}

	for i := range queue1 {
		if queue1[i] != queue2[i] {
			return false
		}
	}
	return true
}

func removeBackspace(queue *[]rune) {
	n := len(*queue)
	writeIndex := 0

	for readIndex := 0; readIndex < n; readIndex++ {
		if (*queue)[readIndex] != '#' {
			(*queue)[writeIndex] = (*queue)[readIndex]
			writeIndex++
		} else {
			if writeIndex > 0 {
				writeIndex--
			}
		}
	}
	*queue = (*queue)[:writeIndex]
}
