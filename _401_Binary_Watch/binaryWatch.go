package _401_Binary_Watch

import (
	"fmt"
)

func BitCount(num int) (c int) {
	for c = 0; num != 0; c++ {
		num = num & (num - 1)
	}

	return c
}

func ReadBinaryWatch(num int) []string {
	var times []string
	for i := 0; i < 12 ; i++ {
		for j := 0; j < 60; j++ {
			tmp := i * 64 + j
			bitCount := BitCount(tmp)

			if num == bitCount {
				time := fmt.Sprintf("%d:%02d", i, j)
				times = append(times, time)
			}
		}

	}

	return  times
}
