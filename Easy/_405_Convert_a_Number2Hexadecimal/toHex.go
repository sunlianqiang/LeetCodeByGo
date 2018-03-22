package _405_Convert_a_Number2Hexadecimal

import "strconv"

func getHexDigit(a int) (hex string){
	if a < 10 {
		return strconv.Itoa(a)
	}

	return string([]byte{'a' + byte(a - 10)})
}
func ToHex(num int) string {
	if 0 == num {
		return "0"
	} else if num < 0 {
		num = 4294967296 + num
	}

	var ret string
	for ; num > 0; {
		a := num % 16
		num = num / 16

		hexDigit := getHexDigit(a)
		ret = hexDigit + ret
	}

	return ret
}