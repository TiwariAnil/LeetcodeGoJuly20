package main

import "fmt"

func prisonAfterNDays(cells []int, N int) []int {

	lenCells := len(cells)
	n :=1
	if n<=N{
		oldVal := cells[0]
		for i:=1; i< lenCells-1; i++{
			if oldVal==cells[i+1]{
				oldVal=cells[i]
				cells[i]=1
			}else {
				oldVal=cells[i]
				cells[i]=0
			}
		}
		cells[0]=0
		cells[lenCells-1]=0
		n++
	}

	N = (N-1)%14+1
	for n:=2; n<=N; n++ {
		oldVal := cells[0]
		for i:=1; i< lenCells-1; i++{
			if oldVal==cells[i+1]{
				oldVal=cells[i]
				cells[i]=1
			}else {
				oldVal=cells[i]
				cells[i]=0
			}
		}
		for _, val := range cells{
			fmt.Print(val)
		}
		fmt.Println()
	}
	return cells
}
func main()  {
	cells := []int{0,0,0,1,1,0,1,0}
	N:=1000000000 //01011010
	N =10000      //00001100
	N =100000     //01000100
	N = 300	//[0,0,0,0,0,1,1,0]
	N = 574
	for _, val := range cells{
		fmt.Print(val)
	}
	fmt.Println()
	res :=prisonAfterNDays(cells, N)
	for _, val := range res{
		fmt.Print(val)
	}
}