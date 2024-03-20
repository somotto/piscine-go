package piscine

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		root = &TreeNode{Data: data}
		return root
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	}
	if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
	}
	if data == root.Data {
		return root
	}
	return root
}
