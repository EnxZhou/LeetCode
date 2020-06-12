package problem0015

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res:=make([][]int,0)
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		if i==0||nums[i]!=nums[i-1]{
			one:=make([]int,3)
			k:=len(nums)-1
			for j:=i+1;j<len(nums)&&j<k;j++{
				if j==i+1||nums[j]!=nums[j-1]{
					for nums[i]+nums[j]+nums[k]>0&&j+1<k{
						k--
					}
					if nums[i]+nums[j]+nums[k]==0{
						one=[]int{nums[i],nums[j],nums[k]}
						res=append(res,one)
					}
				}
			}
		}
	}
	return res
}
