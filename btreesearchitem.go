package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	if elem == root.Data {
		return root
	}
	leftResult := BTreeSearchItem(root.Left, elem)
	if leftResult != nil {
		return leftResult
	}
	rightResult := BTreeSearchItem(root.Right, elem)
	return rightResult
}
