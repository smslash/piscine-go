package piscine

func StrRev(s string) string {
	var arr [1000]byte
	if s[0] == '"' {
		j := 0
		for i := len(s) - 1; i >= 0; i-- {
			arr[j] = s[i]
			j++
		}
		arr[len(s)-1] = '"'
		ans := string(arr[:len(s)])
		return ans
	}
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		arr[j] = s[i]
		j++
	}
	ans := string(arr[:len(s)])
	return ans
}
