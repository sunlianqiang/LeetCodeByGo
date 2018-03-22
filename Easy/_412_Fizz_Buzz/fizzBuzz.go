package _412_Fizz_Buzz

import "strconv"

func FizzBuzz(n int) []string {
	var ret []string
	var i int
	for i = 1; i <= n; i++ {
		if 0 == i % 15 {
			ret = append(ret, "FizzBuzz")
		} else if 0 == i % 3 {
			ret = append(ret, "Fizz")
		} else if 0 == i % 5 {
			ret = append(ret, "Buzz")
		} else {
			ret = append(ret, strconv.Itoa(i))
		}
	}

	return ret
}