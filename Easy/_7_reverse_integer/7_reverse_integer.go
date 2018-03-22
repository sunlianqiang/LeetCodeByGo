package main

import "fmt"

func reverse(x int) int {
	var MIN int = 0x80000000
	var MAX int = 0x7FFFFFFF

	new_int := 0
	var n int
	if x < 0 {
		n = -x
	} else {
		n = x
	}
	for n > 0 {

		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}

	// fmt.Printf("new_int:%+v\n", new_int)
	if new_int > MAX || new_int < -MIN {
		new_int = 0
	}
	if x < 0 {
		new_int = -new_int
	}
	return new_int
}

func main() {
	fmt.Println(reverse(-123456))
	fmt.Println(reverse(0))
	fmt.Println(reverse(1001))
	fmt.Println(reverse(1534236469))

}
