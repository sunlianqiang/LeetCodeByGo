package main

import "fmt"

func twoSum(nums []int, target int) (ret []int) {
	numLen := len(nums)
	for i := 0; i < numLen-1; i++ {
		// fmt.Printf("%d: %d\n", i, nums[i])
		for j := i + 1; j < numLen; j++ {
			sum := nums[i] + nums[j]
			// fmt.Printf("sum:%d = %d + %d\n", sum, nums[i], nums[j])

			if target == sum {
				ret = []int{i, j}

			}
		}
	}
	return ret
}

func main() {
	nums := []int{2, 7, 11, 15}

	target := 9
	ret := twoSum(nums, target)

	fmt.Printf("ret:%+v\n", ret)
}
