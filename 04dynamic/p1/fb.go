package p1

func f1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return f1(n-1) + f1(n-2)
}

func f2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	res, pre, temp := 1, 1, 0
	for i := 3; i <= n; i++ {
		temp = res
		res = pre + res
		pre = temp
	}
	return res
}
