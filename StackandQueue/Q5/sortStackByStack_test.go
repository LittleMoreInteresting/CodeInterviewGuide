package Q5

import (
	"fmt"
	"testing"

	"CodeInterviewGuide/StackandQueue"
)

func TestSortStackByStack(t *testing.T) {
	es := StackandQueue.New()
	es.Push(1)
	es.Push(5)
	es.Push(4)
	es.Push(3)
	SortStackByStack(es)
	fmt.Println(es.Pop())
	fmt.Println(es.Pop())
	fmt.Println(es.Pop())
	fmt.Println(es.Pop())
}
