package main

import "fmt"

func exist(board [][]byte, word string) bool {

	for i:=0; i<len(board); i++{
		for j:=0; j<len(board[0]); j++{
			if len(word)>0 && board[i][j]==word[0] {
				visited := [201][201]uint8{}
				visited[i][j]=1
				val := DFSx(&board, len(board), len(board[0]), i, j, word, 0, &visited)
				if val{
					return true
				}
			}
		}
	}
	return false
}

func DFSx(board *[][]byte, ROW int, COL int,i int, j int, word string, wordIndex int, visited *[201][201]uint8) bool {

	if (*board)[i][j] != word[wordIndex]{
		return false
	}
	if wordIndex == len(word)-1{
		return true
	}
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}
	found :=false
	for k:=0; k<4; k++{
		x:=i+dx[k]
		y:=j+dy[k]
		if x>-1 && y>-1 && x<ROW && y<COL && (*visited)[x][y]==0{
			(*visited)[i][j]=1

			found = DFSx(board, ROW, COL, x, y, word,wordIndex+1, visited)
			(*visited)[i][j]=0
			if found{
				return found
			}

		}
	}
	return false
}

func main() {
	nodes := findOrder(3, [][]int{{0,1},{0,2},{1,2}})
	for _, val := range  nodes{
		fmt.Println(val)
	}
}
