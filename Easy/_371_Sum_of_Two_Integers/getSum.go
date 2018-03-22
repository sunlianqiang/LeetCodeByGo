package _371_Sum_of_Two_Integers

func GetSum(a int, b int) int {
	if a < 0 && b < 0 {
		if a > b {
			a, b = b, a
		}
	} else if a > 0 && b < 0 {
		if a < -b {
			a, b = b, a
		}
	} else if a < 0 && b > 0 {
		if a < b {
			a, b = b, a
		}
	}


	if b > 0 {
		for i := 0; i < b; i++ {
			a++
		}
	} else {
		for i := b; i < 0; i++ {
			a--
		}
	}


	return a
}