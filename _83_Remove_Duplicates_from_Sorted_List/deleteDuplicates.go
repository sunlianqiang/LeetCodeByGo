package _83_Remove_Duplicates_from_Sorted_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}

	newHead := head
	p := head
	q := head.Next
	for ; q != nil ; {
		if p.Val == q.Val {
			p.Next = q.Next
			q.Next = nil
			q = p.Next
		} else {
			p = q
			q = q.Next
		}
	}

	return newHead
}