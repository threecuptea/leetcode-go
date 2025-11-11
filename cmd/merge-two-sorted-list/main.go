package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 21. Merge Two Sorted Lists
* https://leetcode.com/problems/merge-two-sorted-lists/description/
* spicing together with those two list instead creating new nodes
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	dummy := &ListNode{0, nil}
	prev := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next, list1 = list1, list1.Next
		} else {
			prev.Next, list2 = list2, list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}
	return dummy.Next
}
