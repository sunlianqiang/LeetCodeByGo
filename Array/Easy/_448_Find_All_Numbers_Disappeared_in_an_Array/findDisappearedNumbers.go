package _448_Find_All_Numbers_Disappeared_in_an_Array

func FindDisappearedNumbers(nums []int) []int {
	var ret []int
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[i] != nums[nums[i] - 1] {
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
			i--
		}
	}

	for i := 0; i < length; i++ {
		if (i + 1) != nums[i] {
			ret = append(ret, (i+1))
		}
	}

	return ret
}
