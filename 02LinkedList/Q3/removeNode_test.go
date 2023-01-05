package Q3

import (
	"testing"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func Test_removeMidNode(t *testing.T) {
	head := LinkedList.NewBySlice(1)
	head.Show()
	cur := RemoveMidNode(head)
	cur.Show()

	head = LinkedList.NewBySlice(1, 2)
	head.Show()
	cur = RemoveMidNode(head)
	cur.Show()

	head = LinkedList.NewBySlice(1, 2, 3, 4, 5)
	head.Show()
	cur = RemoveMidNode(head)
	cur.Show()
}

func Test_removeMidNode01(t *testing.T) {
	head := LinkedList.NewBySlice(1)
	head.Show()
	cur := RemoveMidNode01(head)
	cur.Show()

	head = LinkedList.NewBySlice(1, 2)
	head.Show()
	cur = RemoveMidNode01(head)
	cur.Show()

	head = LinkedList.NewBySlice(1, 2, 3, 4)
	head.Show()
	cur = RemoveMidNode01(head)
	cur.Show()

	head = LinkedList.NewBySlice(1, 2, 3, 4, 5)
	head.Show()
	cur = RemoveMidNode01(head)
	cur.Show()
}
