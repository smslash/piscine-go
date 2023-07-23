package piscine

func IsNumeric(s string) bool {
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			ans++
		}
	}
	if ans == len(s) {
		return true
	}
	return false
}
