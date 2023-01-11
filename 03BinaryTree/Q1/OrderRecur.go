package Q1

import (
	"fmt"

	BinaryTree "CodeInterviewGuide/03BinaryTree"
)

func ProOrderRecur(node *BinaryTree.TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Value)
	ProOrderRecur(node.Left)
	ProOrderRecur(node.Right)
}

func InOrderRecur(node *BinaryTree.TreeNode) {
	if node == nil {
		return
	}
	ProOrderRecur(node.Left)
	fmt.Printf("%d ", node.Value)
	ProOrderRecur(node.Right)
}

func PosOrderRecur(node *BinaryTree.TreeNode) {
	if node == nil {
		return
	}
	ProOrderRecur(node.Left)
	ProOrderRecur(node.Right)
	fmt.Printf("%d ", node.Value)
}
