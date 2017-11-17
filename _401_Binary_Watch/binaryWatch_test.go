package _401_Binary_Watch

import (
	"testing"
	"fmt"
)

func TestReadBinaryWatch(t *testing.T) {
	var tests = []struct{
		input int
	} {
		{0},
		{1},
	}

	for _, v := range tests {
		ret := ReadBinaryWatch(v.input)

		fmt.Printf("ret:%+v\n", ret)
	}
}
