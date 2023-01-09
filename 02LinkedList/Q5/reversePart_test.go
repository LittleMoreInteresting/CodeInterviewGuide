package Q5

import (
	"testing"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func Test_reverseN(t *testing.T) {
	head := LinkedList.NewBySlice(0, 1, 2, 3, 4, 5)
	head.Show()
	n := ReverseN(head, 3)
	n.Show()
}

func Test_reversePart(t *testing.T) {
	head := LinkedList.NewBySlice(0, 1, 2, 3, 4, 5)
	head.Show()
	n := ReversePart(head, 2, 4)
	n.Show()
}
