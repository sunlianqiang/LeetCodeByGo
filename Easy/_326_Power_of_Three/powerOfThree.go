package _326_Power_of_Three

import (
	"math"
	"fmt"
)

func isPowerOfThree(n int) bool {
	if 0 == n {
		return false
	}
	res := math.Log10(float64(n)) / math.Log10(float64(3))

	var res1 float64
	res1 = res - float64(int(res))
	fmt.Printf("res:%+v, res1:%+v, int res:%+v\n", res, res1, int(res))
	if res1 < 1E-14 {
		return true
	}

	return false
}