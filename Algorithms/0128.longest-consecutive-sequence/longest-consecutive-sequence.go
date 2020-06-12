package problem0128

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums)<=1{
		return len(nums)
	}
	sort.Ints(nums)
	var max,tmp int
	tmp=1
	for i:=0;i<len(nums)-1;i++{
		if nums[i]+1==nums[i+1]{
			tmp++
		}else if nums[i]!=nums[i+1]{
			tmp=1
		}

		if max<tmp{
			max=tmp
		}
	}
	return max
}
