package Q14

import (
	LinkedList "CodeInterviewGuide/02LinkedList"
)

func RemoveValue2(head *LinkedList.Node, num int) *LinkedList.Node {
	dummy := LinkedList.New()
	dummy.Next = head
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == num {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
