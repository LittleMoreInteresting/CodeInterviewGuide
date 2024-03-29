package Q4

import (
	"testing"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func TestReverseList(t *testing.T) {
	head := LinkedList.NewBySlice(1, 2, 3, 4, 5)
	head.Show()
	cur := ReverseList(head)
	cur.Show()
}

func TestReverseListDouble(t *testing.T) {
	head := LinkedList.NewDoubleBySlice(1, 2, 3, 4, 5)
	head.Show()
	cur := ReverseListDouble(head)
	cur.Show()
}
