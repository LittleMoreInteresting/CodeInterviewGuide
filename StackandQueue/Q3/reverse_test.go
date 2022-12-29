package Q3

import (
	"fmt"
	"testing"

	"CodeInterviewGuide/StackandQueue"
)

func TestReverse(t *testing.T) {
	es := StackandQueue.New()
	es.Push(1)
	es.Push(2)
	es.Push(3)
	Reverse(es)
	fmt.Println(es.Pop())
	fmt.Println(es.Pop())
	fmt.Println(es.Pop())

}
