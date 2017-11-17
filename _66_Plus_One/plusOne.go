package _66_Plus_One

func plusOne(digits []int) []int {
	len1 := len(digits)

	digits[len1 - 1] = digits[len1 - 1] + 1
	var carry int
	if digits[len1 - 1] == 10 {
		carry = 1
		digits[len1 - 1] = 0
	}


	for i := len1-2; i >= 0; i-- {
		digits[i] = carry + digits[i]
		if digits[i] == 10 {
			digits[i] = 0
			carry = 1
		} else {
			carry = 0
		}
	}

	if carry == 1 {
		var ret []int
		ret = append([]int{carry}, digits...)
		return ret
	} else {
		return digits
	}
}
