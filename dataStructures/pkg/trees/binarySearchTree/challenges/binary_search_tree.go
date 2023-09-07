package challenges

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key == k {
		return true
	} else if n.Key < k {
		return n.Right.Search(k)
	} else {
		return n.Left.Search(k)
	}
}

func (n *Node) BFS(applyFc func(*Node)) {
	if n == nil {
		return
	}

	queue := []*Node{n}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		applyFc(current)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func (n *Node) InOrderDFS(applyFc func(*Node)) {
	if n == nil {
		return
	}

	n.Left.InOrderDFS(applyFc)
	applyFc(n)
	n.Right.InOrderDFS(applyFc)
}

func (n *Node) PreOrderDFS(applyFc func(*Node)) {
	if n == nil {
		return
	}

	applyFc(n)
	n.Left.PreOrderDFS(applyFc)
	n.Right.PreOrderDFS(applyFc)
}

func (n *Node) PostOrderDFS(applyFc func(*Node)) {
	if n == nil {
		return
	}

	n.Left.PostOrderDFS(applyFc)
	n.Right.PostOrderDFS(applyFc)
	applyFc(n)
}
