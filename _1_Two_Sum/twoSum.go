package _1_Two_Sum

import "fmt"

func TwoSum(nums []int, target int) []int {
	var numMap map[int]int
	numMap = make(map[int]int)
	var ret []int

	for index, val := range nums {
		val2 := target - val
		fmt.Printf("val:%+v, index:%+v, val2:%+v\n", val, index, val2)

		index2, ok := numMap[val2]
		if ok && index != index2 {
			ret = []int{index, index2}
			break
		}

		numMap[val] = index
	}

	fmt.Printf("numMap:%+v\n", numMap)

	return ret
}
