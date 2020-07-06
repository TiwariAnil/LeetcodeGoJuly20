package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getHeightOfTree(root *TreeNode) int {

	if root == nil {
		return 0
	}
	left := 1 + getHeightOfTree(root.Left)
	right := 1 + getHeightOfTree(root.Right)
	return max(left, right)
}

func levelOrderBottomHelp(root *TreeNode, level int, res map[int][]int) {
	if root == nil {
		return
	}
	res[level] = append(res[level], root.Val)
	levelOrderBottomHelp(root.Left, level+1, res)
	levelOrderBottomHelp(root.Right, level+1, res)
}

func levelOrderBottom(root *TreeNode) [][]int {

	var res map[int][]int
	res = make(map[int][]int)

	if root == nil {
		return make([][]int, 0)
	}
	levelOrderBottomHelp(root, 0, res)
	height := getHeightOfTree(root)
	for k, _ := range res {
		if k+1 > height {
			height = k + 1
		}
	}
	result := make([][]int, height)
	for k, v := range res {
		result[height-k-1] = v

	}
	return result
}
