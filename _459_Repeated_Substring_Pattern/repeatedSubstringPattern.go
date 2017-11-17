package _459_Repeated_Substring_Pattern

import "fmt"

func repeatedSubstringPattern(s string) bool {
	lenS := len(s)

	for i := 1; i <= lenS/2; i++ {
		if lenS % i != 0 {
			continue
		}

		var notRepeat bool
		sub := s[:i]
		for j := i; j <= lenS - i; j = j+i {
			if s[j:j+i] != sub {
				notRepeat = true
				fmt.Printf("sub:%+v, sj:%+v\n", sub, s[j:j+i])
				break
			}
		}

		if !notRepeat {
			return true
		}
	}

	return false
}
