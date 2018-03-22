package _551_Student_Attendance_Record_I

func CheckRecord(s string) bool {
	sRune := []rune(s)
	len1 := len(sRune)

	var absentNum int
	for i := 0; i < len1; i++ {
		if sRune[i] == rune('A') {
			absentNum++

			if absentNum >=2 {
				return false
			}
		} else if rune('L') == sRune[i] && i > 0 && i < len1 - 1 && rune('L') == sRune[i-1] && rune('L') == sRune[i+1] {
			return false
		}
	}

	return true
}
