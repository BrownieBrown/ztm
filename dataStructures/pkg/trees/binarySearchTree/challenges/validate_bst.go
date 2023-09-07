package challenges

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	return inOrderTraversal(root, &prev)
}

func inOrderTraversal(node *TreeNode, prev **TreeNode) bool {
	if node == nil {
		return true
	}

	if !inOrderTraversal(node.Left, prev) {
		return false
	}

	if *prev != nil && (*prev).Val >= node.Val {
		return false
	}

	*prev = node
	return inOrderTraversal(node.Right, prev)
}
