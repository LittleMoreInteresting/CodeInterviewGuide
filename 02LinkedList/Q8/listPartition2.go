package Q8

import LinkedList "CodeInterviewGuide/02LinkedList"

func ListPartition2(head *LinkedList.Node, n int) *LinkedList.Node {
	if head == nil {
		return head
	}
	sH := LinkedList.New() // 小于
	sT := sH
	eH := LinkedList.New() // 等于
	eT := eH
	lH := LinkedList.New() // 大于
	lT := lH
	for head != nil {
		switch {
		case head.Val < n:
			sT.Next = head
			sT = sT.Next
		case head.Val == n:
			eT.Next = head
			eT = eT.Next
		case head.Val > n:
			lT.Next = head
			lT = lT.Next
		}
		head = head.Next
	}
	if eH.Next != nil {
		sT.Next = eH.Next
		sT = eT
		sT.Next = nil
	}
	if lH.Next != nil {
		sT.Next = lH.Next
		sT = lT
		sT.Next = nil
	}
	return sH.Next
}
