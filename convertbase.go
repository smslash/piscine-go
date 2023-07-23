package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := 0
	ans := make(map[byte]int)
	for i := 0; i < len(baseFrom); i++ {
		ans[baseFrom[i]] = i
	}
	for i := 0; i < len(nbr); i++ {
		n *= len(baseFrom)
		n += ans[nbr[i]]
	}

	if n == 0 {
		return string(baseTo[0])
	}

	res := []byte{}

	for n > 0 {
		res = append(res, baseTo[n%len(baseTo)])
		n /= len(baseTo)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
