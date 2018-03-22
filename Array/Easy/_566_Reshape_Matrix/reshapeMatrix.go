package _566_Reshape_Matrix

import "fmt"

func MatrixReshape(nums [][]int, r int, c int) [][]int {
	var chanInt chan int
	chanInt = make(chan int)

	var length int
	go func(len *int) {
		for _, v1 := range nums {
			for _, v := range v1 {
				(*len)++
				fmt.Printf("len1:%+v\n", length)
				chanInt <- v
			}
		}

		for {
			chanInt <- 0
		}
	}(&length)


	var ret [][]int
	for i := 0; i < r; i++ {
		var lineRet []int
		for j := 0; j < c; j++ {
			v := <- chanInt

			lineRet = append(lineRet, v)
		}
		ret = append(ret, lineRet)
	}

	fmt.Printf("len2:%+v\n", length)
	fmt.Printf("ret:%+v\n", ret)
	if r * c != length {
		return nums
	}

	return ret
}