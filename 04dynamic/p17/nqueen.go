package p17

import "math"

func Nqueen(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return process1(0, record, n)
}

func process1(i int, record []int, n int) int {
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += process1(i+1, record, n)
		}
	}
	return res
}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if j == record[k] || math.Abs(float64(record[k]-j)) == math.Abs(float64(i-k)) {
			return false
		}
	}
	return true
}
