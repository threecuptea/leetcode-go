package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 19. Remove Nth Node From End of List
* https://leetcode.com/problems/remove-nth-node-from-end-of-list/
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	l, r := dummy, head
	// The idea is to make r -l = n
	i := 1
	for i < n {
		head = head.Next
		i++
	}
	r = head
	if head != nil {
		head = head.Next
		for head != nil {
			// advance l and r
			l, r = l.Next, r.Next
			head = head.Next
		}
	}
	// r should point to the last node and l points to the node before the target node
	l.Next = l.Next.Next
	return dummy.Next
}
