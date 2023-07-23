package piscine

func NRune(s string, n int) rune {
	if len(s) >= n && n > 0 {
		return []rune(s)[n-1]
	} else if n < 0 {
		if len(s) <= n*(-1) {
			return []rune(s)[n-1]
		}
	}
	return 0
}
