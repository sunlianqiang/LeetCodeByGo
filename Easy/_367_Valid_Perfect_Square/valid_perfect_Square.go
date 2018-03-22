package _367_Valid_Perfect_Square

func isPerfectSquare(num int) bool {

	for i := 1; i <= num; i++ {
		if num == i * i {
			return true
		} else if num < i * i {
			return false
		}
	}

	return false
}
