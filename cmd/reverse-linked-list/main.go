package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 206. Reverse Linked List
* https://leetcode.com/problems/reverse-linked-list/description/
 */
func reverseList(head *ListNode) *ListNode {
	// empty or 1
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}
