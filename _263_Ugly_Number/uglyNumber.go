package _263_Ugly_Number

func isUgly(num int) bool {
	if num < 1 {
		return false
	}

	for ; num > 1 && num % 2 == 0 ; {
		num /= 2
	}

	for ; num > 1 && num % 3 == 0 ; {
		num /= 3
	}

	for ; num > 1 && num % 5 == 0 ; {
		num /= 5
	}

	if num == 1 {
		return true
	}

	return false
}
