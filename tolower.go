package piscine

func ToLower(s string) string {
	res := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			v = v + 32
		}
		res += string(v)
	}
	return res
}
