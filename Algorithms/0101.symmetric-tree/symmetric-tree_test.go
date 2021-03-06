package problem0101

import (
	"fmt"
	"leetcode/kit"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0101(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     bool
	}{
		{
			[]int{},
			[]int{},
			true,
		},

		{
			[]int{1, 2, 2},
			[]int{2, 1, 2},
			true,
		},

		{
			[]int{1, 2, 3, 2, 3},
			[]int{2, 3, 1, 2, 3},
			false,
		},
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		// printStruct(kit.PreIn2Tree(tc.pre, tc.in))
		ast.Equal(tc.ans, isSymmetric(kit.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}

func printStruct(root *kit.TreeNode) {
	if root.Left != nil {
		printStruct(root.Left)
	}
	if root.Right != nil {
		printStruct(root.Right)
	}
	fmt.Println(root.Val)
}
