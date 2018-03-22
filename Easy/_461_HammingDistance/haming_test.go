package _461_HammingDistance
import (
	"testing"
)

func Test_HammingDistance(t *testing.T) {
	ret := HammingDistance(1, 4)
	if 2 != ret {
		t.Errorf("test fail, want 2, get %+v\n", ret)
	} else {
		t.Logf("test pass, want 2, get %+v\n", ret)
	}
}