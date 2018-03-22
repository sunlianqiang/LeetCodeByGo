package _258_Add_Digits

func AddDigits(num int) int {
	for ;num >= 10; {
		var sum int
		for ;num >= 10; {
			sum += num % 10
			num = num / 10
		}
		num += sum
	}

	return num
}
