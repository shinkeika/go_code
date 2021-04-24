package main

func recoverTree(root *TreeNode) {
	var nums []int
	// 组装中序遍历
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		nums = append(nums, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	// 查找xy
	x, y := findXY(nums)
	re(root, 2, x, y)
	// 恢复二叉树
}

func findXY(nums []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func re(root *TreeNode, count int, x int, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	re(root.Left, count, x, y)
	re(root.Right, count, x, y)
}
