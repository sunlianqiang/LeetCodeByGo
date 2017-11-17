package _476_Number_Complement


import "testing"


func TestFindComplement(t *testing.T) {
	ret := FindComplement(5)

	if 2 == ret {
		t.Logf("pass")
	}
}
