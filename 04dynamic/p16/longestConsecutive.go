package p16

func LongestConsecutive(nums []int) int {
	mp := make(map[int]int)
	max := 1
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]]; !ok {
			mp[nums[i]] = 1
			if _, lok := mp[nums[i]-1]; lok {
				less := nums[i] - 1
				left := less - mp[less] + 1
				right := nums[i] + mp[nums[i]] - 1
				length := right - left + 1
				mp[left] = length
				mp[right] = length
				max = MaxInt(max, length)
			}
			if _, lok := mp[nums[i]+1]; lok {
				left := nums[i] - mp[nums[i]] + 1
				right := nums[i] + 1 + mp[nums[i]+1] - 1
				length := right - left + 1
				mp[left] = length
				mp[right] = length
				max = MaxInt(max, length)
			}
		}
	}
	return max
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	set := make(map[int]int)
	for idx := range nums {
		set[nums[idx]] = nums[idx]
	}
	res := 0
	for n := range set {
		if _, ok := set[n-1]; ok {
			continue
		}
		curNum := n
		curLen := 1
		for _, ok := set[curNum+1]; ok; _, ok = set[curNum+1] {
			curNum += 1
			curLen += 1
			res = MaxInt(res, curLen)
		}
	}
	return res
}
