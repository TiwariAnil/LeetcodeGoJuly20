package main

import (
	"fmt"
	"sort"
)
func InsertSorted(s []int, e int) []int {
	i := sort.Search(len(s), func(i int) bool { return s[i] > e })
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}
func nthUglyNumber(n int) int {
	N := n+2
	var pool []int
	var mp map[int]bool
	mp = make(map[int]bool)
	pool = append(pool, 1)
	for i := 0; i < N; i++ {
		val := pool[i]
		valN := val*2
		if _, ok:= mp[valN]; !ok {
			mp[valN]=true
			pool = InsertSorted(pool, valN)
		}

		valN = val*3
		if _, ok:= mp[valN]; !ok {
			mp[valN]=true
			pool = InsertSorted(pool, valN)
		}

		valN = val*5
		if _, ok:= mp[valN]; !ok {
			mp[valN]=true
			pool = InsertSorted(pool, valN)
		}
	}
	return pool[n-1]
}
func main() {
	fmt.Println(nthUglyNumber(1690))
	fmt.Println(nthUglyNumber(10))

}
