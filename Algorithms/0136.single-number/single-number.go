package problem0136

func singleNumber(nums []int) int {
	var a int
	for i:=0;i<len(nums);i++ {
		a=a^nums[i]
	}
	return a
}
