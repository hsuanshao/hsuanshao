package main

func main() {}

// Runtime: 0 ms
// Memory Usage: 2.3 MB

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	var nodes []*ListNode
	for {
		nodes = append(nodes, head)
		if head.Next == nil {
			break
		}
		head = head.Next
		count++
	}
	removed := count - n + 1

	for i := count; i >= 0; i-- {
		head = nodes[i]

		if i == removed {
			var prevNode, newNode *ListNode
			if i+1 < count+1 {
				newNode = nodes[i+1]
			}

			if i-1 >= 0 {
				prevNode = nodes[i-1]
			}
			head = updateNode(prevNode, newNode)
		}
	}

	return head
}

func updateNode(node *ListNode, newNode *ListNode) *ListNode {
	if node == nil && newNode == nil {
		return nil
	}

	if node == nil && newNode != nil {
		return newNode
	}

	node.Next = newNode
	return node
}
