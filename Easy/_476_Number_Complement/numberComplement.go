package _476_Number_Complement

func FindComplement(num int) int {
	tmpnum := num
	var ret int
	for tmpnum > 0 {
		ret = ret * 2 + 1
		tmpnum = tmpnum >> 1
	}

	ret = ret ^ num

	return ret
}