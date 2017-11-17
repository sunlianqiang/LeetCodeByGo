package _202_Happy_Number

func getNewNum(num int) (newNum int) {
	for ; num > 0 ; {
		a := num % 10
		num = num / 10
		newNum += a * a
	}

	return newNum
}
func isHappy(n int) bool {
	var numMap map[int]bool
	numMap = make(map[int]bool)
	var newNum int
	numMap[n] = true
	newNum = n

	for {
		if 1 == newNum {
			return true
		}
		newNum = getNewNum(newNum)

		_, ok := numMap[newNum]
		if ok {
			return false
		} else {
			numMap[newNum] = true
		}
	}
}
