package problem1300

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	arr   []int
	target   int
	ans int
}{

	{
		[]int{4,9,3},
		10,
		3,
	},
	{
		[]int{2,3,5},
		10,
		5,
	},
	{
		[]int{60864,25176,27249,21296,20204},
		56803,
		11361,
	},

	// 可以有多个 testcase
}

func Test_findBestValue(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findBestValue(tc.arr, tc.target), "输入:%v", tc)
	}
}

func Benchmark_findBestValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findBestValue(tc.arr, tc.target)
		}
	}
}
