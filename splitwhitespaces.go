package piscine

func CheckSymbols(s string) bool {
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '\n' || s[i] != '\t' || s[i] != ' ' {
			ans++
		}
	}
	if ans != 0 {
		return true
	}
	return false
}

func SplitWhiteSpaces(s string) []string {
	s += string(' ')
	str := []string{}
	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\t' || s[i] == ' ' {
			if CheckSymbols(ans) {
				str = append(str, ans)
			}
			ans = ""
			continue
		}
		ans += string(s[i])
	}
	return str
}
