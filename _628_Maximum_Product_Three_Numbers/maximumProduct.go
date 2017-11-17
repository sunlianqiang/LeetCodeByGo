package _628_Maximum_Product_Three_Numbers

import (
	"sort"
	"fmt"
)

func MaximumProduct(nums []int) int {
	len1 := len(nums)
	if len1 == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	sort.Ints(nums)
	fmt.Printf("int:%+v\n", nums)

	if (nums[0] * nums[1]) > (nums[len1 - 2] * nums[len1 - 3]) {
		return nums[0] * nums[1] * nums[len1 - 1]
	} else {
		return nums[len1 - 1] * nums[len1 - 2] * nums[len1 - 3]
	}

}
