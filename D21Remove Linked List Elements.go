package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	var prev *ListNode
	var root *ListNode
	prev = nil
	for head != nil {
		if head.Val != val {
			if prev == nil {
				root = head
				prev = head
			} else {
				prev.Next = head
				prev = head
			}
		} else {
			if prev != nil {
				prev.Next = nil
			}
		}
		head = head.Next
	}
	return root
}
