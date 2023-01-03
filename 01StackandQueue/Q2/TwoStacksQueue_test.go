package Q2

import (
	"fmt"
	"testing"
)

func TestTwoStacksQueue(t *testing.T) {
	ts := New()
	ts.Add(1)
	ts.Add(2)
	ts.Add(3)
	ts.Add(4)
	fmt.Println(ts.Poll())
	fmt.Println(ts.Poll())
	fmt.Println(ts.Poll())
	fmt.Println(ts.Peek())
}
