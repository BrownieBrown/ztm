package challenges

// Node Definition
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	// base case
	if root == nil {
		return root
	}

	current := root
	for current != nil {
		if current.Child != nil {
			// store next node before we overwrite it
			next := current.Next

			// recursively flatten the child list
			child := flatten(current.Child)

			// connect current node to child
			current.Next = child
			child.Prev = current

			// connect the tail of the child list back to the next node
			// traverse till the end of the flattened child list
			for child.Next != nil {
				child = child.Next
			}

			child.Next = next
			if next != nil {
				next.Prev = child
			}

			// clear the child's pointer since its flattened
			current.Child = nil

			// continue from the next node
			current = current.Next

		} else {
			current = current.Next
		}
	}

	return root
}
