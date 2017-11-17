package _21_Merge_Two_Sorted_Lists

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}

	var p3 *ListNode
	p1, p2 := l1, l2
	if p1.Val < p2.Val {
		p3 = p1
		p1 = p1.Next
		p3.Next = nil
	} else {
		p3 = p2
		p2 = p2.Next
		p3.Next = nil
	}
	l3 := p3

	for ;p1 != nil && p2 != nil; {
		if p1.Val < p2.Val {
			tmp := p1
			p1 = p1.Next
			p3.Next = tmp
			p3 = p3.Next
			p3.Next = nil
		} else {
			tmp := p2
			p2 = p2.Next
			p3.Next = tmp
			p3 = p3.Next
			p3.Next = nil
		}
	}

	if p1 != nil {
		p3.Next = p1
	} else {
		p3.Next = p2
	}
	return l3
}
