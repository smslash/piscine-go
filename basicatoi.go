package piscine

func Add(a byte, len int) int {
	ans := 1
	if a == '1' {
		ans = 1
		for i := 1; i <= len; i++ {
			ans *= 10
		}
		return ans
	} else if a == '2' {
		ans = 2
		for i := 1; i <= len; i++ {
			ans *= 10
		}
		return ans
	} else if a == '3' {
		ans = 3
		for i := 1; i <= len; i++ {
			ans *= 10
		}
		return ans
	} else if a == '4' {
		ans = 4
		for i := 1; i <= len; i++ {
			ans *= 10
		}
		return ans
	} else if a == '5' {
		ans = 5
		for i := 1; i <= len; i++ {
			ans *= 10
		}
		return ans
	} else if a == '6' {
		ans = 6
		for i := 1; i <= len; i++ {
			ans *= 10
		}
		return ans
	} else if a == '7' {
		ans = 7
		for i := 1; i <= len; i++ {
			ans *= 10
		}
		return ans
	} else if a == '8' {
		ans = 8
		for i := 1; i <= len; i++ {
			ans *= 10
		}
		return ans
	} else if a == '9' {
		ans = 9
		for i := 1; i <= len; i++ {
			ans *= 10
		}
		return ans
	}
	return 0
}

func BasicAtoi(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += Add(s[i], len(s)-i-1)
	}
	return ans
}
