package linkedLists

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

// printListData Method
func (l *LinkedList) printListData() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

// deleteWithValue Method
func (l *LinkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	current := l.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
		l.length--
	}
}

// prepend Method
func (l *LinkedList) prepend(n *node) {
	n.next = l.head
	l.head = n
	l.length++
}

// append Method
func (l *LinkedList) append(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	l.length++
}

// search Method
func (l *LinkedList) search(value int) *node {
	current := l.head
	for current != nil {
		if current.data == value {
			return current
		}
		current = current.next
	}
	return nil
}

// reverse Method
func (l *LinkedList) reverse() {
	var prev *node
	current := l.head
	var next *node
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}
