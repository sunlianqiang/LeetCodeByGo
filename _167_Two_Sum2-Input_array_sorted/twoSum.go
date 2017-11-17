package _167_Two_Sum2_Input_array_sorted

func TwoSum(numbers []int, target int) []int {
	var ret []int
	len1 := len(numbers)

	for i, j := 0, len1 - 1; i < j; {
		if target == numbers[i] + numbers[j] {
			i++
			j++
			ret = []int{i, j}
			break
		} else if target > numbers[i] + numbers[j] {
			i++
		} else if target < numbers[i] + numbers[j] {
			j--
		}
	}

	return ret
}