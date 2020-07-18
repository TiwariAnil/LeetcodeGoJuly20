package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	key int
	val int
}
func topKFrequent(nums []int, k int) []int {

	var mp map[int]int
	mp = make(map[int]int)
	for _,val := range nums {
		if v,ok := mp[val]; ok{
			mp[val] = v+1
		}else{
			mp[val]=1
		}
	}
	var list []Pair
	for k,v := range mp {
		list = append(list, Pair{k,v})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].val > list[j].val
	})
	var res []int
	for i:=0; i<k; i++{
		res = append(res, list[i].key)
	}
	return res
}
func main() {
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3,3,3,3}, 2))
}
