package main

import "fmt"

//为了不占内存不能转成字符串反转后比较
// 为了防止溢出，每次比较最高位和最低位
func isPalindrome(x int) bool {
	fmt.Printf("x:%+v\n", x)
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	//取最高位
	high := 10

	for x/high > 9 {
		high *= 10
	}

	for x > 0 {
		fmt.Printf("new_x:%+v, high:%+v\n", x, high)
		numHigh := x / high
		numLow := x % 10
		fmt.Printf("numHigh:%d, numLow:%d\n", numHigh, numLow)
		if numHigh != numLow {
			return false
		}

		x = x % high
		x /= 10
		high /= 100
	}

	return true
}

func main() {

	fmt.Printf("isPalindrome:%+v\n\n", isPalindrome(123))
	fmt.Printf("isPalindrome:%+v\n\n", isPalindrome(1001))
	fmt.Printf("isPalindrome:%+v\n\n", isPalindrome(1000021))
	fmt.Printf("isPalindrome:%+v\n\n", isPalindrome(434))
	fmt.Printf("isPalindrome:%+v\n\n", isPalindrome(4334))
	fmt.Printf("isPalindrome:%+v\n\n", isPalindrome(-4334))
	fmt.Printf("isPalindrome:%+v\n\n", isPalindrome(0))
}
