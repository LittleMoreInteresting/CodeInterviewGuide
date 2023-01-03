package Q2

import (
	"CodeInterviewGuide/StackandQueue"
)

/**
编写一个类，用两个栈实现队列，支持队列的基本操作（add、poll、peek）
*/

type TwoStacksQueue struct {
	stackPush *StackandQueue.EasyStack
	stackPop  *StackandQueue.EasyStack
}

func New() *TwoStacksQueue {
	return &TwoStacksQueue{
		stackPush: StackandQueue.New(),
		stackPop:  StackandQueue.New(),
	}
}

func (ts *TwoStacksQueue) Add(n int) {
	ts.stackPush.Push(n)
}

func (ts *TwoStacksQueue) Poll() int {
	if ts.stackPop.Empty() && ts.stackPush.Empty() {
		panic("Queue is empty")
	}
	if ts.stackPop.Empty() {
		// Pop 栈为空时移动Push 栈数据
		for !ts.stackPush.Empty() {
			ts.stackPop.Push(ts.stackPush.Pop())
		}
	}
	return ts.stackPop.Pop()
}

func (ts *TwoStacksQueue) Peek() int {
	if ts.stackPop.Empty() && ts.stackPush.Empty() {
		panic("Queue is empty")
	}
	if ts.stackPop.Empty() {
		// Pop 栈为空时移动Push 栈数据
		for !ts.stackPush.Empty() {
			ts.stackPop.Push(ts.stackPush.Pop())
		}
	}
	return ts.stackPop.Peek()
}
