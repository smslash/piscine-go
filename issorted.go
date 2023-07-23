package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) > 1 {
		if a[0] > a[1] {
			for i := 0; i < len(a)-1; i++ {
				ans := f(a[i], a[i+1])
				if ans < 0 {
					return false
				}
			}
			return true
		} else {
			for i := 0; i < len(a)-1; i++ {
				ans := f(a[i], a[i+1])
				if ans > 0 {
					return false
				}
			}
			return true
		}
	}
	return true
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}
