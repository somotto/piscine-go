package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	return BTreeSearchItem(root.Right, elem)
}
