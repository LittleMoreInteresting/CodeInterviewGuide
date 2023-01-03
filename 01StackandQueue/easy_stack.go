package StackandQueue

type EasyStack struct {
	num []int
}

func New() *EasyStack {
	return &EasyStack{
		num: []int{},
	}
}

func (es *EasyStack) Pop() (p int) {
	l := len(es.num)
	if l == 0 {
		panic("stack is empty")
	}
	p = es.num[l-1]
	es.num = es.num[:l-1]
	return
}

func (es *EasyStack) Push(n int) {
	es.num = append(es.num, n)
}

func (es *EasyStack) Peek() int {
	return es.num[len(es.num)-1]
}

func (es *EasyStack) Empty() bool {
	return len(es.num) == 0
}
