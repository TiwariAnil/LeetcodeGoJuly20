package main

import "fmt"

func singleNumber(nums []int) []int {

	A :=0
	B :=0
	xor :=0
	for _,val := range nums{
		xor = xor ^ val
	}
	// xor is now A ^ B

	firstChangeBit := (xor & (xor-1)) ^ xor
	for _,val := range nums{
		if (firstChangeBit & val)==0 {
			A = A ^ val
		}else{
			B = B ^ val
		}
	}
	return []int{A,B}
}
func main() {
}
