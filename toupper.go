package piscine

func ToUpper(s string) string {
	result := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			v = v - 32
		}
		result += string(v)
	}
	return result
}
