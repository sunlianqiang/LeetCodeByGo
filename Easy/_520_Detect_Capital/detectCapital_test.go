package _520_Detect_Capital

import (
	"testing"
)

func TestDetectCapitalUse(t *testing.T) {
	var tests = []struct {
		input   string
		output bool
	}{
		{"USA", true},
		{"FlaG", false},
		{"Leetcode", true},
	}
	for _, test := range tests {
		ret := DetectCapitalUse(test.input)
		if ret == test.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v", ret, test.output)
		}

	}
}