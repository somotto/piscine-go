package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data == elem {
		return root
	}
	if root.Left != nil {
		return BTreeSearchItem(root.Left, elem)
	}
	if root.Right != nil {
		return BTreeSearchItem(root.Right, elem)
	}
	return nil
}
