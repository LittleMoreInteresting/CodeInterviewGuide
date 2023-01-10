package Q8

import (
	"testing"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func Test_listPartition2(t *testing.T) {
	head := LinkedList.NewBySlice(9, 0, 4, 3, 5, 3, 1)
	head.Show()
	ret := ListPartition2(head, 3)
	ret.Show()
}
