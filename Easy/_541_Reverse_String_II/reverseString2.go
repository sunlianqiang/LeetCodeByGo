package _541_Reverse_String_II

func reverseRune(r []rune, start, end int) {
	for ; start < end; 	{
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
}

func reverseStr(s string, k int) string {
	sRune := []rune(s)
	len1 := len(s)
	for i := 0; i < len1; i += k * 2 {
		end := i + k - 1
		if end > len1 - 1{
			end = len1 - 1
		}

		reverseRune(sRune, i, end)
	}

	ret := string(sRune)

	return ret
}
