package piscine

func IsLower(s string) bool {
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			ans++
		}
	}
	if ans == len(s) {
		return true
	}
	return false
}
