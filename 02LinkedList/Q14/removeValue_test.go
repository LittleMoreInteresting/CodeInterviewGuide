package Q14

import (
	"testing"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func Test_removeValue2(t *testing.T) {
	head := LinkedList.NewBySlice(1, 2, 3, 2, 4)
	head.Show()
	res := RemoveValue2(head, 2)
	res.Show()
}
