package piscine

func Concat(str1 string, str2 string) string {
	s := append([]rune(str1)[:], []rune(str2)...)
	return string(s)
}
