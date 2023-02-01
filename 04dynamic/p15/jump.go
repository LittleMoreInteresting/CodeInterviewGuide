package p15

func Jump(nums []int) int {
	lenght := len(nums)
	if lenght == 0 {
		return 0
	}
	dp := 0
	cur := 0
	next := 0
	for i := 0; i < lenght; i++ {
		if cur < i {
			dp++
			cur = next
		}
		next = MaxInt(next, i+nums[i])
	}
	return dp
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 3，2，3，1，1，4
