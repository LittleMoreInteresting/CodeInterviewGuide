package Q5

import LinkedList "CodeInterviewGuide/02LinkedList"

//https://labuladong.gitee.io/algo/di-yi-zhan-da78c/shou-ba-sh-8f30d/di-gui-mo--1eaae/
//采用递归+拆分 思想实现

//# 01 翻转前N个节点
var successor *LinkedList.Node

func ReverseN(node *LinkedList.Node, n int) *LinkedList.Node {
	if n == 1 {
		successor = node.Next
		return node
	}
	last := ReverseN(node.Next, n-1)
	node.Next.Next = node
	node.Next = successor
	return last
}

//# 02 结合第一步方法 + 递归
func ReversePart(node *LinkedList.Node, from, to int) *LinkedList.Node {
	if from == 1 {
		return ReverseN(node, to)
	}
	node.Next = ReversePart(node.Next, from-1, to-1)
	return node
}
