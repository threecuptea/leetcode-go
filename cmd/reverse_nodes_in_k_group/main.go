package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 25. Reverse Nodes in k-Group
* https://leetcode.com/problems/reverse-nodes-in-k-group/description/
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	dummy := &ListNode{0, head}
	var curr, lp *ListNode = head, dummy
	for curr != nil {
		var tmp, res *ListNode = curr, nil
		i, partial := 0, true
		// Look ahead to make sure it is not partial.  We won't reverse partial group
		for tmp != nil && i < k {
			// store the first pointer
			if res == nil {
				res = tmp
			}
			i++
			tmp = tmp.Next
		}
		if i == k {
			partial = false
		}
		if partial {
			lp.Next = res
			break
		}
		i = 0
		var newLp, prev *ListNode = curr, nil
		for i < k {
			i++
			next := curr.Next
			curr.Next = prev
			prev, curr = curr, next
		}
		// prev points to the last in k group, curr points to the following node
		lp.Next, lp = prev, newLp
	}
	return dummy.Next
}
