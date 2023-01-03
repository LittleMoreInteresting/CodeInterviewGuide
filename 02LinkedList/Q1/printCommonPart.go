package Q1

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func New() *Node {
	return &Node{}
}

func printCommonPart(head1, head2 *Node) {
	fmt.Println("Common Part: ")
	for head1 != nil && head2 != nil {
		if head1.Val == head2.Val {
			fmt.Printf("%d ", head1.Val)
			head1 = head1.Next
			head2 = head2.Next
			continue
		}
		if head1.Val > head2.Val {
			head2 = head2.Next
		} else if head1.Val < head2.Val {
			head1 = head1.Next
		}
	}
}
