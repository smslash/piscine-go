package piscine

func lowwer(s string) string {
	res := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			v = v + 32
		}
		res += string(v)
	}
	return res
}

func upper(s byte) byte {
	return s - 32
}

func Capitalize(s string) string {
	s = lowwer(s)
	res := ""
	if s[0] >= 'a' && s[0] <= 'z' {
		res += string(upper(s[0]))
	} else if s[0] == 34 {
		res += string(92)
		res += string(34)
	} else if s[0] == 92 {
		res += string(92)
	} else {
		res += string(s[0])
	}
	for i := 1; i < len(s); i++ {
		if s[i] == 34 {
			res += string(92)
			res += string(34)
		} else if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			if (s[i-1] < 'a' || s[i-1] > 'z') && (s[i-1] < '0' || s[i-1] > '9') {
				if s[i] >= 'a' && s[i] <= 'z' {
					res += string(upper(s[i]))
				} else {
					res += string(s[i])
				}
			} else {
				res += string(s[i])
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}
