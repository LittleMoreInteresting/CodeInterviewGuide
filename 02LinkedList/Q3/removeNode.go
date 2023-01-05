package Q3

import LinkedList "CodeInterviewGuide/02LinkedList"

func RemoveMidNode(head *LinkedList.Node) *LinkedList.Node {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		return head.Next
	}
	pre, cur := head, head.Next.Next
	for cur.Next != nil && cur.Next.Next != nil {
		pre = pre.Next
		cur = cur.Next.Next
	}
	pre.Next = pre.Next.Next
	return head
}

func RemoveMidNode01(head *LinkedList.Node) *LinkedList.Node {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := LinkedList.New()
	dummy.Next = head
	fast, slow := dummy.Next.Next, dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}
