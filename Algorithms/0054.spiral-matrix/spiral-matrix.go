package problem0054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}
	var left, right, top, bottom int
	left = 0
	right = len(matrix[0]) - 1
	top = 0
	bottom = len(matrix) - 1
	var res []int
	for left <= right && top <= bottom {
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++
		if bottom < top {
			break
		}
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for j := right; j >= left; j-- {
			res = append(res, matrix[bottom][j])
		}
		bottom--
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if right < left {
			break
		}
	}
	return res
}
