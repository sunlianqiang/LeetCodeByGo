package _551_Student_Attendance_Record_I

import "testing"

func TestCheckRecord(t *testing.T) {
	var tests = []struct{
		input string
		output bool
	} {
		{"PPALLL", false},
	}

	for _, v := range tests {
		ret := CheckRecord(v.input)

		if ret == v.output {
			t.Logf("pass")
		} else {
			t.Errorf("fail, want %+v, get %+v\n", v.output, ret)
		}
	}
}
