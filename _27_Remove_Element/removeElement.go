package _27_Remove_Element

func removeElement(nums []int, val int) int {
	len1 := len(nums)
	var len2 int
	for i := 0; i < len1; i++{
		if val != nums[i] {
			if len2 < i {
				nums[len2] = nums[i]
			}
			len2++
		}
	}
	nums = nums[:len2]
	return len2
}
