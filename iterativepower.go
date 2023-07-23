package piscine

func IterativePower(nb int, power int) int {
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
