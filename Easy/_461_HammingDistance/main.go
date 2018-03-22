package _461_HammingDistance

func HammingWeight(z int) (w int){

	for z > 0 {
		tmp := z % 2
		if 1 == tmp {
			w++
		}
		z = z /2
	}

	return w
}
func HammingDistance(x int, y int) int {

	z := x ^ y

	return HammingWeight(z)
}