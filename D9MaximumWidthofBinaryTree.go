package main

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

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

func levelOrderBottomHelp(root *TreeNode, level int, value int, maxD *map[int]int, minD *map[int]int) {
	if root == nil {
		return
	}
	if val, ok := (*maxD)[level]; ok{
		if val < value {
			(*maxD)[level]=val
		}
	}else{
		(*maxD)[level]=val
	}
	if val, ok := (*minD)[level]; ok{
		if val > value {
			(*minD)[level]=val
		}
	}else{
		(*minD)[level]=val
	}
	levelOrderBottomHelp(root.Left, level+1, 2*value - 1, maxD, minD)
	levelOrderBottomHelp(root.Right, level+1, value + 1, maxD, minD)
}

func widthOfBinaryTree(root *TreeNode) int {

	var maxD map[int]int
	maxD = make(map[int]int)

	var minD map[int]int
	minD = make(map[int]int)

	if root == nil {
		return 0
	}
	levelOrderBottomHelp(root, 0, 0, &maxD, &minD)

	dia :=0
	for k, maxVal := range maxD {
		if minVal, ok := minD[k]; ok {
			if dia < (maxVal - minVal){
				dia = maxVal-minVal
			}
		}
	}

	return dia
}
