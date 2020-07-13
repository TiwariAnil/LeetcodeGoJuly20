package main

import "fmt"

func reverseBits(num uint32) uint32 {

	var temp uint32
	temp = num
	var ans uint32
	var flag uint32
	ans = 0
	for num>0 {
		flag = 0
		if num % 2==1 {
			flag = 1
		}
		num = num >> 1
		ans = ans + flag
		if num > 0{
			ans = ans << 1
		}
	}
	//ans = ans >> 1
	bit :=0
	for temp>0 {
		bit++
		temp = temp>>1
	}
	//fmt.Println(ans)
	//fmt.Println(bit)
	for bit<32 {
		ans = ans << 1
		bit++
	}
	return ans
}

func main()  {

	fmt.Println(reverseBits(4294967293)) //322 122 547 1
	fmt.Println(reverseBits(43261596)) //964176192
x :=1073741823
x = x << 1
x = x << 1
fmt.Println(x)
}