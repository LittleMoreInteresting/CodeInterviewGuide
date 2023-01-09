package LinkedList

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func New() *Node {
	return &Node{}
}

func NewBySlice(nums ...int) *Node {
	head := New()
	cur := head
	for i := range nums {
		next := New()
		next.Val = nums[i]
		cur.Next = next
		cur = next
	}
	return head.Next
}

func NewRoundBySlice(nums ...int) *Node {
	head := New()
	cur := head
	var next *Node
	for i := range nums {
		next = New()
		next.Val = nums[i]
		cur.Next = next
		cur = next
	}
	next.Next = head.Next
	return head.Next
}

func (n *Node) Show() {
	fmt.Print("start:")
	p := n
	for p != nil {
		fmt.Printf("->%v ", p.Val)
		p = p.Next
		if p == n {
			fmt.Printf("=>%v ", p.Val)
			return
		}
	}
	fmt.Println()
}

type DoubleNode struct {
	Val  int
	Prev *DoubleNode
	Next *DoubleNode
}

func NewDoubleNode() *DoubleNode {
	return &DoubleNode{}
}

func NewDoubleBySlice(nums ...int) *DoubleNode {
	head := NewDoubleNode()
	cur := head
	for i := range nums {
		next := NewDoubleNode()
		next.Val = nums[i]
		cur.Next = next
		next.Prev = cur
		cur = next
	}
	return head.Next
}

func (n *DoubleNode) Show() {
	for n != nil {
		fmt.Printf("<==>%v ", n.Val)
		n = n.Next
	}
	fmt.Println()
}
