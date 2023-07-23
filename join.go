package piscine

func Join(strs []string, sep string) string {
	res := ""
	sum := 0
	for _, v := range strs {
		res += v
		sum++
		if sum == len(strs) {
			return res
		}
		res += sep
	}
	return res
}
