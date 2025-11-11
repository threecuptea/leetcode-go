package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 82. Remove Duplicates from Sorted List II
* https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prev, curr := dummy, head
	for curr != nil {
		cleanRes := false
		for curr.Next != nil && curr.Next.Val == curr.Val {
			prev.Next, curr = curr.Next, curr.Next
			cleanRes = true
		}
		if cleanRes {
			// need to clean up the residual current
			prev.Next, curr = curr.Next, curr.Next
		} else {
			prev, curr = curr, curr.Next
		}
	}
	return dummy.Next
}
