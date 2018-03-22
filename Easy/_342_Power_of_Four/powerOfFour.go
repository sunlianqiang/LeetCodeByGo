package _342_Power_of_Four

import (
	"math"
	"fmt"
)

func isPowerOfFour(num int) bool {
	if 0 == num {
		return false
	}
	res := math.Log10(float64(num)) / math.Log10(float64(4))

	var res1 float64
	res1 = res - float64(int(res))
	fmt.Printf("res:%+v, res1:%+v, int res:%+v\n", res, res1, int(res))
	if res1 < 1E-14 {
		return true
	}

	return false
}
