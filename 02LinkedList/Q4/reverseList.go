package Q4

import LinkedList "CodeInterviewGuide/02LinkedList"

func ReverseList(head *LinkedList.Node) *LinkedList.Node {
	var pre, next *LinkedList.Node
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
