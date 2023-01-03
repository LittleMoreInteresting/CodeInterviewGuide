package Q3

import "CodeInterviewGuide/StackandQueue"

//将栈stack的栈底元素返回并移除
func getAndRemoveLastElement(es *StackandQueue.EasyStack) int {
	res := es.Pop()
	if es.Empty() {
		return res
	}
	last := getAndRemoveLastElement(es)
	es.Push(res)
	return last
}

func Reverse(es *StackandQueue.EasyStack) {
	if es.Empty() {
		return
	}
	last := getAndRemoveLastElement(es)
	Reverse(es)
	es.Push(last)
}
