package main

import (
	"fmt"
	"sort"
)

func subRecur(nums []int, index int, list []int, res *[][]int) {

	if index == len(nums) {
		*res = append(*res, list)
		return
	}
	subRecur(nums, index+1, list, res)
	subRecur(nums, index+1, append(list, nums[index]), res)
}

func subsets(nums []int) [][]int {

	var res [][]int
	var list []int
	if len(nums) == 0 {
		return res
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	subRecur(nums, 0, list, &res)
	return res
}

func main() {

	list := []int{9, 0, 3, 5, 7}

	res := subsets(list)
	for _, ar := range res {
		for _, j := range ar {
			fmt.Print(j)
			fmt.Print(", ")
		}
		fmt.Println()
	}

}
