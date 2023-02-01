package p14

func Win(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	return MaxInt(first(nums, 0, length-1), second(nums, 0, length-1))
}

//先手
func first(nums []int, l, r int) int {
	if l == r {
		return nums[l]
	}
	return MaxInt(nums[l]+second(nums, l+1, r), nums[r]+second(nums, l, r-1))
}

//后手
func second(nums []int, l, r int) int {
	if l == r {
		return 0
	}
	return MinInt(first(nums, l+1, r), first(nums, l, r-1))
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
