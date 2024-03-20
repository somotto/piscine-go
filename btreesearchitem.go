package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data == elem {
		return root
	}
	leftResult := BTreeSearchItem(root.Left, elem)
	if leftResult != nil {
		return leftResult
	}
	return BTreeSearchItem(root.Right, elem)
}
