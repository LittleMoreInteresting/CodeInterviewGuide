package BinaryTree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func New() *TreeNode {
	return &TreeNode{}
}
