package piscine

func AppendRange(min, max int) []int {
	s := []int(nil)
	if max <= min {
		return s
	}
	for i := min; i < max; i++ {
		s = append(s, i)
	}
	return s
}
