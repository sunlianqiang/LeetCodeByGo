package _521_Longest_Uncommon_Subsequence_I

func FindLUSlength(a string, b string) int {
	lena := len(a)
	lenb := len(b)

	if lena > lenb {
		return lena
	} else if lena < lenb {
		return lenb
	}else if a != b {
		return lena
	} else {
		return -1
	}
}