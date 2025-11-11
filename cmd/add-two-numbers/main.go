package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 2. Add Two Numbers
* https://leetcode.com/problems/add-two-numbers/
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// At least 1 in each list
	val := l1.Val + l2.Val
	head := &ListNode{val % 10, nil}
	result, l1, l2 := head, l1.Next, l2.Next
	val = val / 10
	for l1 != nil || l2 != nil {
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		result.Next = &ListNode{val % 10, nil}
		result = result.Next
		val = val / 10 // carry over from the previous round
	}
	// Don't forget the final carry over
	if val == 1 {
		result.Next = &ListNode{1, nil}
	}
	return head
}

func buildListNodes(vals []int) *ListNode {
	var head *ListNode = nil
	n := len(vals)
	for i := n - 1; i >= 0; i-- {
		head = &ListNode{vals[i], head}
	}
	return head
}

func convertToSlice(head *ListNode) []int {
	var lst []int
	for head != nil {
		lst = append(lst, head.Val)
		head = head.Next
	}

	return lst
}

func main() {
	ints1 := []int{2, 4, 3}
	ints2 := []int{5, 6, 4}
	fmt.Printf("The inputs are %v, %v\n", ints1, ints2)
	l1, l2 := buildListNodes(ints1), buildListNodes(ints2)
	h := addTwoNumbers(l1, l2)
	fmt.Printf("The output is %v\n", convertToSlice(h))
	fmt.Println()

	ints1 = []int{9, 9, 9, 9, 9, 9, 9}
	ints2 = []int{9, 9, 9, 9}
	fmt.Printf("The inputs are %v, %v\n", ints1, ints2)
	l1, l2 = buildListNodes(ints1), buildListNodes(ints2)
	h = addTwoNumbers(l1, l2)
	fmt.Printf("The output is %v\n", convertToSlice(h))

}
