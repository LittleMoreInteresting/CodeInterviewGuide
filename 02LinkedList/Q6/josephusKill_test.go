package Q6

import (
	"testing"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func Test_josephusKill1(t *testing.T) {
	head := LinkedList.NewRoundBySlice(1, 2, 3, 4, 5)
	head.Show()
	/*l := JosephusKill1(head, 3)
	l.Show()*/
	l2 := JosephusKill2(head, 3)
	l2.Show()
}
