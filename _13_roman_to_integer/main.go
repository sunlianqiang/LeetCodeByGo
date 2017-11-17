package main

import "fmt"

// Ⅰ（1）、Ⅴ（5）、Ⅹ（10）、Ⅼ（50）、Ⅽ（100）、Ⅾ（500）和Ⅿ（1000）

func romanToInt(s string) int {
	var romanMap map[string]int
	romanMap = make(map[string]int)
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000

	var sum int
	sLen := len(s)
	sum = romanMap[s[sLen-1:sLen]]
	for i := sLen - 2; i >= 0; i-- {
		s1 := s[i+1 : i+2]
		s2 := s[i : i+1]
		num1 := romanMap[s1]
		num2 := romanMap[s2]
		fmt.Printf("num1:%+v, num2:%+v\n", num1, num2)

		if num1 <= num2 {
			sum = sum + num2
		} else {
			sum = sum - num2
		}

		fmt.Printf("%+v, sum:%+v\n", i, sum)

	}

	fmt.Printf("***********\nroman:%+v, integer:%+v\n*******\n", s, sum)
	return sum
}

func main() {

	romanToInt("I")
	romanToInt("X")
	romanToInt("C")
	romanToInt("D")
	romanToInt("DCXXI")

}
