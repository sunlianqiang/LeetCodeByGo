package _35_Search_Insert_Position

func SearchInsert(nums []int, target int) int {

	len1 := len(nums)
	if target > nums[len1-1] {
		return len1
	}
	l := 0
	r := len1
	for ; l <= r ; {
		middle := (l + r) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			l = middle + 1
		} else {
			r = middle - 1
		}
	}

	// l == r
	return r + 1
}
