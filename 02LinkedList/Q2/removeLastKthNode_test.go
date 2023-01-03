package Q2

import (
	"testing"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func Test_RemoveLastKthNode(t *testing.T) {
	head := LinkedList.NewBySlice(1, 2, 3, 4, 5)
	head.Show()
	cur := RemoveLastKthNode(head, 6)
	cur.Show()
}
