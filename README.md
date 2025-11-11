# leetcode-go

This is a repo that I showcase my capabity in Go Programming with Leetcode problems, including

* **Linked List Cycle**: Use Floyd's Cycle-Finding Algorithm.  It is much faster than 'out of box' implementation.
* **Add Two Numbers**: Use 'val' double duty to sum up carrier, l1.val and l2.val and also as the carrier for the next round. Make sure to record the last carrier even if both l1 and l2 are exhausted.
* **LRU cache**: Design  a **[Least Recently Used (LRU) cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU)**.
  The implmentation is using double linked list nodes backing LRU cache and add left and right dummy nodes to point to the least used node and the recently used node.  That is the common implmentation.
* **Reverse linked list and reverse linked list II**:  Follow 3-step reverse process. The reverse linked list II do need to fast forward to the left position then reverse (right-left+1) time then adjust point. It's 3-stage process.
* **Remove Nth Node From End of List and Rotate List**: I group them together because both require the left and right pointer. The left point is at the position left to the target node or the last node after the rotation. Need to use the left and/ or the right to re-assign pointer.  Also the head is subject to change.  Therefore, using dummy node to code around edge cases.  It's possible to rotate more than the length of list nodes. I get the net modulo to avoid NullPointer issue.
* **Reverse Nodes in k-Group**: It is considered the toughest one in linked list group because

  1) Cannot and should not reverse the residual list which is shorter than k.  The residual partial list need to be handled differently.
  2) Reverse multiple k-group.  That requires save and reset lp (the last node of the previous k-group or the dummy node)

All Go modules got accepted by Leetcode and passed all tests. I added testcases for those Go modules commited initially.  Will spare time to add testcases for the rest of Go modules.  More are coming up...
