package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	l, r := dummy, head
	i := 0
	// I have to loop through it to get around nil pointer issue.
	for head != nil {
		i++
		head = head.Next
	}
	// i is the length here
	k, i, head = k%i, 1, r
	if k == 0 {
		return head
	}
	for i < k {
		head = head.Next
		i++
	}
	// l, r has k distance
	r = head
	if head != nil {
		head = head.Next
		for head != nil {
			// advance l and r
			l, r = l.Next, r.Next
			head = head.Next
		}
	}
	// r should point to the last node and l points to the last node after the rotation.
	r.Next = dummy.Next
	dummy.Next = l.Next
	l.Next = nil
	return dummy.Next
}
