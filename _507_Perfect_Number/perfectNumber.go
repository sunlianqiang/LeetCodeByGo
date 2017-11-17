package _507_Perfect_Number

import "math"

func CheckPerfectNumber(num int) bool {
	if 1 == num {
		return false
	}
	sqrt1 := int(math.Sqrt(float64(num)))
	var sum int
	sum = 1
	for i := 2; i <= sqrt1; i++ {
		if num % i == 0 {
			sum += i
			sum += num /i
		}

	}

	if sum == num {
		return true
	}

	return false
}
