package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}

	arr := map[rune]bool{}
	for _, r := range base {
		if r == '+' || r == '-' || arr[r] {
			return 0
		}
		arr[r] = true
	}

	ans := len(base)
	res := 0
	for _, r := range s {
		n := -1
		for i, b := range base {
			if r == b {
				n = i
				break
			}
		}
		if n == -1 {
			return 0
		}
		res = res*ans + n
	}

	return res
}
