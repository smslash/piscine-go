package piscine

func IterativeFactorial(nb int) int {
	if nb >= 2 && nb <= 20 {
		ans := 1
		for i := 2; i <= nb; i++ {
			ans *= i
			if ans > 9223372036854775807 {
				return 0
			}
		}
		return ans
	} else if nb == 0 || nb == 1 {
		return 1
	}
	return 0
}
