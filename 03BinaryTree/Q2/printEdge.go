package Q2

import (
	"fmt"

	BinaryTree "CodeInterviewGuide/03BinaryTree"
)

var edgeMap [][]*BinaryTree.TreeNode

func PrintEdge1(head *BinaryTree.TreeNode) {
	if head == nil {
		return
	}
	height := getHeight(head, 0)

	edgeMap = make([][]*BinaryTree.TreeNode, height)
	for i := range edgeMap {
		edgeMap[i] = make([]*BinaryTree.TreeNode, 2)
	}
	setEdgeMap(head, 0)
	for i := range edgeMap {
		fmt.Printf("%d ", edgeMap[i][0].Value)
	}
	printLeafNotInMap(head, 0)
	for i := range edgeMap {
		if edgeMap[i][0] == edgeMap[i][1] {
			continue
		}
		fmt.Printf("%d ", edgeMap[i][1].Value)
	}
}
func setEdgeMap(head *BinaryTree.TreeNode, l int) {
	if head == nil {
		return
	}
	if edgeMap[l][0] == nil {
		edgeMap[l][0] = head
	}
	edgeMap[l][1] = head
	setEdgeMap(head.Left, l+1)
	setEdgeMap(head.Right, l+1)
}

func printLeafNotInMap(head *BinaryTree.TreeNode, l int) {
	if head == nil {
		return
	}
	if head.Left == nil && head.Right == nil &&
		edgeMap[l][0] != head && edgeMap[l][1] != head {
		fmt.Printf("%d ", head.Value)
	}
}

func getHeight(head *BinaryTree.TreeNode, h int) int {
	if head == nil {
		return h
	}
	return Max(getHeight(head.Left, h+1), getHeight(head.Right, h+1))
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PrintEdge2(head *BinaryTree.TreeNode) {
	if head == nil {
		return
	}
	fmt.Printf("%d ", head.Value)
	if head.Left != nil && head.Right != nil {
		printLeftEdge(head.Left, true)
		printRightEdge(head.Left, true)
	} else {
		var next *BinaryTree.TreeNode
		if head.Left != nil {
			next = head.Left
		} else {
			next = head.Right
		}
		PrintEdge2(next)
	}
}

func printLeftEdge(head *BinaryTree.TreeNode, print bool) {
	if head == nil {
		return
	}
	if print || (head.Left == nil && head.Right == nil) {
		fmt.Printf("%d ", head.Value)
	}
	printLeftEdge(head.Left, print)
	printLeftEdge(head.Right, print && head.Left == nil)
}

func printRightEdge(head *BinaryTree.TreeNode, print bool) {
	if head == nil {
		return
	}
	printRightEdge(head.Left, print && head.Right == nil)
	printRightEdge(head.Right, print)
	if print || (head.Left == nil && head.Right == nil) {
		fmt.Printf("%d ", head.Value)
	}
}
