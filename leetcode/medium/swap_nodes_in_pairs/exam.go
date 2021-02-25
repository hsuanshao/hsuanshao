package main

// Runtime: 0 ms
// Memory Usage: 2.1 MB

func main() {}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	head.Next = swapPairs(res.Next)
	res.Next = head
	return res
}
