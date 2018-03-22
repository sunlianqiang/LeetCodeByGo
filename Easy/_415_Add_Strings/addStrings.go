package _415_Add_Strings

import (
	"strconv"
)


func AddStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	var ret string
	var carry int
	var digit int

	for i, j := len1, len2; i > 0 || j > 0 || carry > 0; {
		digit = carry
		if i > 0 {
			i--
			digit += int(num1[i] - '0')
		}
		if j > 0 {
			j--
			digit += int(num2[j] - '0')
		}

		if digit > 9 {
			carry = 1
		} else {
			carry = 0
		}

		ret = strconv.Itoa(digit % 10) + ret
	}

	return ret
}
