package Q1

import (
	"testing"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func Test_PrintCommonPart(t *testing.T) {
	head1 := LinkedList.NewBySlice(1, 2, 3, 4, 5)
	head2 := LinkedList.NewBySlice(3, 4, 5)
	head1.Show()
	head2.Show()
	PrintCommonPart(head1, head2)

}
