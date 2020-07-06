package main

import "fmt"

func hammingDistance(x int, y int) int {

	cnt :=0
	for x>0 || y>0 {
		if x%2 != y%2 {
			cnt++
		}
		x = x>>1
		y = y>>1
	}
	return cnt
}
func main()  {

	fmt.Println(hammingDistance(1, 4))
	var val int
	val = 100000
	fmt.Println(val)
	val = val >> 1
	fmt.Println(val)
	val = val >> 1
	fmt.Println(val)
	val = val >> 1
	fmt.Println(val)
	val = val >> 1
	fmt.Println(val)

}