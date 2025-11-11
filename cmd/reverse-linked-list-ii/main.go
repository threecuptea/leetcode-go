package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 92. Reverse Linked List II
* https://leetcode.com/problems/reverse-linked-list-ii/description/
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{0, head}
	prev, curr := dummy, head
	i := 1
	for i < left {
		prev, curr = curr, curr.Next
		i++
	}
	var lp *ListNode
	lp, prev = prev, nil
	i = left
	for i <= right {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
		i++
	}
	lp.Next.Next = curr
	lp.Next = prev
	return dummy.Next
}
