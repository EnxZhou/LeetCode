package problem1300

import (
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	sum:=0
	for i:=0;i<len(arr);i++{
		aa:=(target-sum)/(len(arr)-i)
		tt:=float64(target-sum)/float64(len(arr)-i)
		if arr[i]>=aa{
			if (tt-float64(aa))>0.5{
				return aa+1
			}
			return aa
		}
		sum+=arr[i]
	}
	return arr[len(arr)-1]
}
