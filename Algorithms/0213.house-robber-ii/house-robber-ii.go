package problem0213

func rob(nums []int) int {
	size := len(nums)
	switch size {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	return max(robbing(nums[1:]), robbing(nums[:size-1]))
}

func robbing(nums []int) int {
	nm1, nm2, result := 0, 0, 0
	for _, v := range nums {
		result = max(nm2+v, nm1)
		nm2 = nm1
		nm1 = result
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
