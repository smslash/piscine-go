package piscine

func RecursiveFactorial(nb int) int {
	if nb >= 2 && nb <= 20 {
		return nb * IterativeFactorial(nb-1)
	}

	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb == 1 {
		return 1
	}
	return 0
}
