package Q6

import (
	LinkedList "CodeInterviewGuide/02LinkedList"
)

// O(n*m)
func JosephusKill1(head *LinkedList.Node, m int) *LinkedList.Node {
	if head == nil || head.Next == head || m < 1 {
		return head
	}
	last := head
	for last.Next != head {
		last = last.Next
	}
	count := 0
	for head != last {
		count++
		if count == m {
			last.Next = head.Next
			count = 0
		} else {
			last = last.Next
		}
		head = last.Next
	}
	return head
}

func JosephusKill2(head *LinkedList.Node, m int) *LinkedList.Node {
	if head == nil || head.Next == head || m < 1 {
		return head
	}
	cur := head.Next
	tmp := 1
	for cur != head {
		tmp++
		cur = cur.Next
	}
	tmp = getLive(tmp, m)
	for tmp--; tmp != 0; tmp-- {
		head = head.Next
	}
	head.Next = head
	return head
}

func getLive(i, m int) int {
	if i == 1 {
		return 1
	}
	return (getLive(i-1, m)+m-1)%i + 1
}
