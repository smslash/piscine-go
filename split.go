package piscine

func Split(s, sep string) []string {
	var res []string
	for {
		ans := index(s, sep)
		if ans == -1 {
			res = append(res, s)
			break
		}
		res = append(res, s[:ans])
		s = s[ans+len(sep):]
	}
	return res
}

func index(s, sep string) int {
	n := len(sep)
	for i := 0; i <= len(s)-n; i++ {
		if s[i:i+n] == sep {
			return i
		}
	}
	return -1
}
