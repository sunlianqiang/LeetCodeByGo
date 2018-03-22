package _412_Fizz_Buzz

import "testing"

func arrayStringEqual(want, ret []string) bool {
	len1 := len(want)
	len2 := len(ret)

	if len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if want[i] != ret[i] {
			return false
		}
	}

	return true
}
func TestFizzBuzz(t *testing.T) {
	input := 15
	want := []string{"1",
					 "2",
					 "Fizz",
					 "4",
					 "Buzz",
					 "Fizz",
					 "7",
					 "8",
					 "Fizz",
					 "Buzz",
					 "11",
					 "Fizz",
					 "13",
					 "14",
					 "FizzBuzz"}

	ret := FizzBuzz(input)

	equal := arrayStringEqual(want, ret)

	if equal {
		t.Logf("pass")
	} else {
		t.Errorf("fail, want %+v, get %+v", want, ret)
	}
}
