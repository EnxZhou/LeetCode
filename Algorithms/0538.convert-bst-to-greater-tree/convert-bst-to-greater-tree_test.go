package problem0538

import (
	"fmt"
	"testing"

	"leetcode/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ansPost []int
}{

	{
		[]int{5, 2, 13},
		[]int{2, 5, 13},
		[]int{20, 13, 18},
	},
	{
		[]int{2},
		[]int{2},
		[]int{2},
	},
	{
		[]int{5,2},
		[]int{2,5},
		[]int{7,5},
	},
	{
		[]int{1,0,-2,4,3},
		[]int{-2,0,1,3,4},
		[]int{6,8,7,4,8},
	},

	// 可以有多个 testcase
}

func Test_convertBST(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ans := kit.Tree2Postorder(convertBST(root))
		ast.Equal(tc.ansPost, ans)
	}
}

func Benchmark_convertBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			convertBST(root)
		}
	}
}
