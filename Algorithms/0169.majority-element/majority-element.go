package problem0169

func majorityElement1(nums []int) int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = 1
		} else {
			m[nums[i]] += 1
		}
	}
	for i, v := range m {
		if v > len(nums)/2 {
			return i
		}
	}
	return 0
}

func majorityElement(nums []int) int {
	count := 1
	current := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == current {
			count++
		} else {
			count--
			if count == 0 {
				current = nums[i+1]
			}
		}
	}
	return current
}
