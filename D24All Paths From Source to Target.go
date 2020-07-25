package main

import "fmt"

func DFSPath(graph *[][]int, visited []int, node int, path []int, allPath *[][]int, finalNode int) {

	if node == finalNode {
		var temp []int
		for _,val := range path{
			temp = append(temp, val)
		}
		*allPath = append(*allPath, temp)
		return
	}
	connectionToStartNode := (*graph)[node]
	for _, nextNode := range connectionToStartNode {
		if visited[nextNode] == 0 {
			visited[nextNode] = 1
			temp := append(path, nextNode)
			DFSPath(graph, visited, nextNode, temp, allPath, finalNode)
			visited[nextNode] = 0
		}
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {

	var visited []int
	visited = append(visited, 1)
	for i := 0; i < 100; i++ {
		visited = append(visited, 0)
	}
	var allPath [][]int
	visited[0]=1
	DFSPath(&graph, visited, 0, []int{0}, &allPath, len(graph)-1)
	return allPath
}


func main() {
	//fmt.Println(allPathsSourceTarget([][]int{{1,2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{3,1},{4,6,7,2,5},{4,6,3},{6,4},{7,6,5},{6},{7},{}}))
}
