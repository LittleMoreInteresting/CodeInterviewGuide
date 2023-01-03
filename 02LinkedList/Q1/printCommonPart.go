package Q1

import (
	"fmt"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func PrintCommonPart(head1, head2 *LinkedList.Node) {
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
