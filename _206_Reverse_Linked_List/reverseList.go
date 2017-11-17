package _206_Reverse_Linked_List

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

func reverseList(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}

	ret := head
	head = head.Next
	ret.Next = nil

	for ; head != nil; {
		tmp := head
		head = head.Next
		tmp.Next = ret

		ret = tmp
	}

	return ret
}
