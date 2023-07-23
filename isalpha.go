package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			ans++
		}
	}
	if ans == len(s) {
		return true
	}
	return false
}
