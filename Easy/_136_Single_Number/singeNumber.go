package _136_Single_Number

func singleNumber(nums []int) int {
	var ret int
	ret = 0
	for _,v := range nums {
		ret = ret ^ v
	}
	return ret
}
