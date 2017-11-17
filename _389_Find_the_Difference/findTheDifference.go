package _389_Find_the_Difference

func FindTheDifference(s string, t string) byte {
	var ret byte
	for _, v := range s {
		ret = ret ^ byte(v)
	}

	for _, v := range t {
		ret = ret ^ byte(v)
	}
	return ret
}
