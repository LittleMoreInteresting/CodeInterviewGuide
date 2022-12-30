package Q7

func getMaxWindow(arr []int, w int) []int {
	l := len(arr)
	if l == 0 || w < 0 || l < w {
		return []int{}
	}
	qmax := []int{}
	res := make([]int, l-w+1)
	idx := 0
	for i := range arr {

		for len(qmax) > 0 && arr[qmax[len(qmax)-1]] <= arr[i] {
			qmax = qmax[:len(qmax)-1]
		}
		qmax = append(qmax, i)
		if qmax[0] == i-w {
			qmax = qmax[1:]
		}
		if i >= w-1 {
			res[idx] = arr[qmax[0]]
			idx++
		}
	}
	return res
}
