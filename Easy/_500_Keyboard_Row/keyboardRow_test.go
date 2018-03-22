package _500_Keyboard_Row

import "testing"

func TestFindWords(t *testing.T) {

	words := []string{
		"Hello", "Alaska", "Dad", "Peace"}
	ret := FindWords(words)

	wanted := map[string]int{
		"Alaska": 1,
		"Dad": 2,
	}

	for _, v := range ret {
		_, ok := wanted[v]

		if !ok {
			t.Errorf("fail, not want but have %+v\n", v)
			return
		}
	}

	t.Logf("pass")
}
