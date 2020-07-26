package main

import "fmt"


func getMin(nums []int ,st int, en int) int {

	if en==st {
		return nums[st]
	}
	if en-st==1{
		if nums[st] < nums[en] {
			return nums[st]
		}
		return nums[en]
	}
	mid := st + (en-st)/2
	if nums[mid]==nums[en] && nums[mid]==nums[st]{
		vL := getMin(nums, st, mid)
		vR := getMin(nums, mid, en)
		if vL < vR{
			return vL
		}
		return vR
	}
	if nums[mid] <= nums[en] && nums[mid] >= nums[st]{
		return getMin(nums, st, mid)
	}
	if nums[mid] <= nums[st] && nums[mid] <=nums[en]{
		return getMin(nums, st, mid)
	}
	return getMin(nums, mid, en)
}

func findMin(nums []int) int {

	return getMin(nums, 0, len(nums)-1);
}

func main() {
	fmt.Println(findMin([]int{6,6, 6, 6, 6, 6, 6, 6,6 ,6 ,6, 6,6 ,6,6 ,6, 8, 8, 8, 8, 1,3,5, 6, 6, 6, 6}))
}
