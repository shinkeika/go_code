package main

import "math"

func maxPathSum(root *TreeNode) int {
	retMat := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		innerMax := left + right + root.Val
		retMat = max(innerMax, retMat)
		outMax := root.Val + max(left, right)
		return max(outMax, 0)
	}
	dfs(root)
	return retMat
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
