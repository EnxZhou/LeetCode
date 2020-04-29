package problem0198

func rob(nums []int) int {
	nm1, nm2, result := 0, 0, 0
	for _, v := range nums {
		result = max(nm2+v, nm1)
		nm2 = nm1
		nm1 = result
	}
	return result
}

func rob1(nums []int) int {
	a, b := 0, 0
	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(b+v, a)
		}
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
