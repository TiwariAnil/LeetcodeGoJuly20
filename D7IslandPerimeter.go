package main

import "fmt"

func islandPerimeter(grid [][]int) int {

	rows := len(grid)
	col := len(grid[0])
	wall:=0
	for i, row := range grid{
		for j,_ := range row{
			if grid[i][j]==1 {
				wall = wall + getWaterTouches(i, j, rows, col, &grid)
			}
		}
	}
	return wall
}

func getWaterTouches(x, y, row, col int, grid *[][]int) int {

	touch :=0
	_x := []int{-1, 0, 1, 0}
	_y := []int{0, -1, 0, 1}
	for i:=0; i<4; i++ {
		x__ := x+ _x[i]
		y__ := y+ _y[i]

		touch = touch + getValid(x__, y__, &row, &col, grid)
	}
	return touch
}
func getValid(x, y int, row* int, col* int, grid *[][]int) int {
	if x < 0 {
		return 1;
	}
	if x >= *row {
		return 1
	}
	if y< 0{
		return 1
	}
	if y >= *col{
		return 1
	}
	if (*grid)[x][y]==0{
		return 1
	}
	return 0
}
func main()  {
	 matrix :=[][]int{
		{0,0,1},
		{0,1,1},
		 {0,0,0},
		 {0,0,0},
	}
	fmt.Println(islandPerimeter(matrix))
}