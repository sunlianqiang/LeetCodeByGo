package _83_Remove_Duplicates_from_Sorted_List

import "testing"

func InitList(nums []int) (head *ListNode)  {
	len1 := len(nums)
	if len1 == 0 {
		return nil
	}

	head = new(ListNode)
	head.Val = nums[0]
	head.Next = nil

	p := head
	for i := 1; i < len1; i++ {
		q := new(ListNode)
		q.Val = nums[i]
		q.Next = nil

		p.Next = q
		p = q
	}

	return head
}

func ListEqual(list1, list2 *ListNode) bool {
	if nil == list1 && nil == list2 {
		return true
	} else if nil != list1 && nil == list2 {
		return false
	} else if nil == list1 && nil != list2 {
		return false
	}

	for ; nil != list1 && nil != list2; {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}

	if nil == list1 && nil == list2 {
		return true
	}

	return false
}

func TestDeleteDuplicates(t *testing.T) {
	input1 := []int{1, 1, 2, 3, 3}
	inputList := InitList(input1)
	output1 := []int{1, 2, 3}
	outputList := InitList(output1)
	var tests = []struct{
		input *ListNode
		output *ListNode
	} {
		{inputList, outputList},
	}

	for _, v := range tests {
		ret := DeleteDuplicates(v.input)
		ok := ListEqual(ret, v.output)
		if ok {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v\n", v.output, ret)
		}
	}
}
