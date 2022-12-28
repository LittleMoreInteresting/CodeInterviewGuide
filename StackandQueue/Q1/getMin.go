package Q1

/**
设计一个有getMin功能的栈
【题目】

实现一个特殊的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作。

【要求】

1.pop、push、getMin操作的时间复杂度都是O (1)。

2．设计的栈类型可以使用现成的栈结构。
*/

type MinStack struct {
	data []int // 存放栈数据
	min  []int // 存放栈最小值
}

func New() *MinStack {
	return &MinStack{
		data: []int{},
		min:  []int{},
	}
}

func (ms *MinStack) Push(num int) {
	ms.data = append(ms.data, num)
	ml := len(ms.min)
	if ml == 0 || num <= ms.GetMin() {
		ms.min = append(ms.min, num)
	}
}

func (ms *MinStack) Pop() (num int) {
	dl, ml := len(ms.data), len(ms.min)
	num = ms.data[dl-1]
	if num == ms.GetMin() {
		ms.min = ms.min[:ml-1]
	}
	ms.data = ms.data[:dl-1]
	return
}

func (ms *MinStack) GetMin() int {

	return ms.min[len(ms.min)-1]
}
