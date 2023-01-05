package Q2

import (
	"fmt"

	LinkedList "CodeInterviewGuide/02LinkedList"
)

func RemoveLastKthNode(head *LinkedList.Node, k int) *LinkedList.Node {
	if head == nil || k < 1 {
		return head
	}
	cur := head
	for cur != nil {
		k--
		cur = cur.Next
	}
	if k > 0 {
		return head
	}
	if k == 0 {
		return head.Next
	}
	cur = head
	k++
	for k != 0 {
		fmt.Println(k)
		cur = cur.Next
		k++
	}
	cur.Next = cur.Next.Next
	return head
}

//快慢指针
func RemoveLastKthNode01(head *LinkedList.Node, k int) *LinkedList.Node {
	fast := head
	for i := 0; i < k; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return head
}

func RemoveLastKthDoubleNode(head *LinkedList.DoubleNode, k int) *LinkedList.DoubleNode {

	return head
}
