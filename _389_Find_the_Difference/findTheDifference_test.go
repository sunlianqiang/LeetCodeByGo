package _389_Find_the_Difference

import "testing"

func TestFindTheDifference(t *testing.T) {
	var tests = []struct{
		s string
		t string
		output byte
	}{
		{"abcd","abcde", byte('e')},
	}

	for _, test := range tests {
		ret := FindTheDifference(test.s, test.t)

		if ret == test.output {

		}
	}

}
