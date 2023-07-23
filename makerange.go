package piscine

func MakeRange(min, max int) []int {
	s := []int(nil)
	if max <= min {
		return s
	}

	ans := make([]int, max-min)

	k := 0
	for i := min; i < max; i++ {
		ans[k] = i
		k++
	}

	return ans
}
