package Q1

import (
	"fmt"
	"testing"
)

func TestMinStack_GetMin(t *testing.T) {
	ms := New()
	ms.Push(3)
	fmt.Println("min", ms.GetMin())
	ms.Push(4)
	fmt.Println("min", ms.GetMin())
	ms.Push(5)
	fmt.Println("min", ms.GetMin())
	ms.Push(1)
	fmt.Println("min", ms.GetMin())
	ms.Push(2)
	fmt.Println("min", ms.GetMin())
	ms.Push(1)
	fmt.Println("min", ms.GetMin())
	fmt.Println("POP", ms.Pop())
	fmt.Println("min", ms.GetMin())
	fmt.Println("POP", ms.Pop())
	fmt.Println("min", ms.GetMin())
	fmt.Println("POP", ms.Pop())
	fmt.Println("min", ms.GetMin())
}
