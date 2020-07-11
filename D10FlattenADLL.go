package main

type Node struct {
     Val int
     Prev *Node
     Next *Node
     Child *Node
 }


/* TODO recursive soln
func keepAdding(root *Node, curr *Node) *Node{

	if root.Next == nil && root.Child == nil {
		return nil
	}
	child := root.Child
	next := root.Next
	if child != nil {

		temp := Node{root.Child.Val, curr, nil, nil}
		curr.Next = &temp
		curr = keepAdding(root.Child, &temp)
	}
	if next!= nil{
		temp := Node{next.Val, curr, nil, nil}
		curr.Next = &temp
		curr = keepAdding(root.Next, curr)
	}
	return curr
}*/

func DFS(root *Node, list *[]Node) {

	if root == nil {
		return
	}
	(*list) = append((*list), *root)
	DFS( root.Child, list)
	DFS( root.Next, list)
}



func flatten(root *Node) *Node {

	var list []Node
	DFS(root, &list)
	var curr *Node
	if list!=nil && len(list)>0{
		curr=&list[0]
		root = curr
	}
	for i:=1; i<len(list); i++ {
		curr.Next = &list[i]
		list[i].Prev = curr
		curr = &list[i]
	}
	return root
}