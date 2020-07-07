package main

func plusOne(digits []int) []int {

	var num []int

	carry :=1
	for i:=len(digits)-1; i>-1; i--{

		val:=digits[i]+carry
		if(val==10){
			carry=1
			num = append([]int{0}, num...)
		}else{
			carry=0
			num = append([]int{val}, num...)
		}
	}
	if carry==1{
		num = append([]int{carry}, num...)
	}
	return num
}

func main()  {

}