package _504_Base7

import "strconv"

func ConvertToBase7(num int) string {
	var ret string
	if num < 0 {
		ret += "-"
		num = -num
	}

	var base7 string
	if 0 == num {
		base7 = "0"
	} else {
		for ;num > 0; {
			base7 = strconv.Itoa(num % 7) + base7
			num = num / 7
		}
	}

	ret += base7

	return ret
}
