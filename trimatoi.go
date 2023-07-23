package piscine

func Ppower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 1 {
		return nb
	}
	ans := 1
	for i := 1; i <= power; i++ {
		ans = ans * nb
	}
	return ans
}

func Checker(s byte) int {
	if s == '0' {
		return 0
	} else if s == '1' {
		return 1
	} else if s == '2' {
		return 2
	} else if s == '3' {
		return 3
	} else if s == '4' {
		return 4
	} else if s == '5' {
		return 5
	} else if s == '6' {
		return 6
	} else if s == '7' {
		return 7
	} else if s == '8' {
		return 8
	} else {
		return 9
	}
}

func TrimAtoi(s string) int {
	res := []int{}
	b := 1
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			break
		}
		if s[i] == '-' {
			b = -1
			break
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			res = append(res, Checker(s[i]))
		}
	}
	for i := 0; i < len(res); i++ {
		ans = ans + res[i]*Ppower(10, len(res)-i-1)
	}
	return ans * b
}
