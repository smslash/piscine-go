package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = n * (-1)
	}
	res := []int{}
	for i := n; i > 0; i /= 10 {
		res = append(res, i%10)
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res); j++ {
			if res[i] > res[j] {
				tmp := res[i]
				res[i] = res[j]
				res[j] = tmp
			}
		}
	}

	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(rune(res[i]) + '0')
	}
}
