package Q7

import (
	"fmt"
	"testing"
)

func Test_getMaxWindow(t *testing.T) {
	arr := []int{4, 3, 5, 4, 3, 3, 6, 7}
	res := getMaxWindow(arr, 3)
	fmt.Println(res) //5,5,5,4,6,7
}
