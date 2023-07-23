package piscine

func Abort(a, b, c, d, e int) int {
	res := []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return res[2]
}
