package piscine

func Rot14(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			if s[i]+14 > 122 {
				res += string(s[i] + 14 - 26)
			} else {
				res += string(s[i] + 14)
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			if s[i]+14 > 90 {
				res += string(s[i] + 14 - 26)
			} else {
				res += string(s[i] + 14)
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}
