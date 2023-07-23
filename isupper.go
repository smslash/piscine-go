package piscine

func IsUpper(s string) bool {
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			ans++
		}
	}
	if ans == len(s) {
		return true
	}
	return false
}
