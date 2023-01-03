package Q5

import "CodeInterviewGuide/StackandQueue"

func SortStackByStack(stack *StackandQueue.EasyStack) {
	help := StackandQueue.New()
	for !stack.Empty() {
		cur := stack.Pop()
		for !help.Empty() && help.Peek() < cur {
			stack.Push(help.Pop())
		}
		help.Push(cur)
	}
	for !help.Empty() {
		stack.Push(help.Pop())
	}
}
