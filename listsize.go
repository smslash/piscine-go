package piscine

func ListSize(l *List) int {
	it := l.Head
	ans := 0
	for it != nil {
		ans++
		it = it.Next
	}
	return ans
}
