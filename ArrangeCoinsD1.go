package main

import (
	"math"
)

func arrangeCoins(n int) int {
	root := math.Sqrt(float64(n))
	for true {
		if x := sumOfFirstN(int(root)); x > n {
			root--
			break
		} else {
			root++
		}
	}
	return int(root)
}

func sumOfFirstN(n int) int {
	return n * (n + 1) / 2
}
func main() {

}
