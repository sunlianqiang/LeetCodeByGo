package _434_Number_of_Segments_in_String

func countSegments(s string) int {
	len1 := len(s)
	if 0 == len1 {
		return 0
	}

	var count int
	for i := 0; i < len1; {
		for ;i < len1 && ' ' == s[i];i++ {}

		if i >= len1 {
			break
		}
		count++
		for ;i < len1 && ' ' != s[i]; i++ {}
	}

	return count
}
