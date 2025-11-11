package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
*  - 141. Linked List Cycle
*    https://leetcode.com/problems/linked-list-cycle/description/
*    Use Floyd's Cycle-Finding Algorithm.  It is much faster than 'out of box' implementation
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	vals := []int{3, 2, 0, -4}
	pos, n := 1, len(vals)
	var head, last *ListNode = nil, nil
	for i := n - 1; i >= 0; i-- {
		head = &ListNode{vals[i], head}
		if i == n-1 {
			last = head
		}
		if i == pos {
			last.Next = head
		}
	}
	fmt.Printf("The Linked List Input, %v, pos=%d\n", vals, pos)
	fmt.Println(hasCycle(head))
}
