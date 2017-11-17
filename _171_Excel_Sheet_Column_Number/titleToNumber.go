package _171_Excel_Sheet_Column_Number

func toNumber(tmp byte) int {
	return int(tmp - byte('A')) + 1
}
func TitleToNumber(s string) int {
	var ret int

	len1 := len(s)
	for i := 0; i < len1; i++ {
		tmp := s[i]
		tmpNum := toNumber(tmp)
		ret = ret * 26 + tmpNum
	}

	return ret
}