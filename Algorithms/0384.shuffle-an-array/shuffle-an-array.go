package problem0384

import (
	"math/rand"
)

// Solution 是答案
type Solution struct {
	A []int
	B []int
}

func Constructor(nums []int) Solution {
	a := make([]int, len(nums))
	b := make([]int, len(nums))
	copy(a, nums)
	copy(b, nums)
	return Solution{a, b}
}

func (this *Solution) Reset() []int {
	this.A = this.B
	return this.A
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.A); i++ {
		a := rand.Int()
		this.A[i], this.A[a%(i+1)] = this.A[a%(i+1)], this.A[i]
	}
	return this.A
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
