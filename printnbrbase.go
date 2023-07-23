package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	res := ""
	if nbr < 0 {
		z01.PrintRune('-')
		var unbr uint = uint(-nbr)
		var ans uint = uint(len(base))
		for unbr >= ans {
			res = string(base[unbr%ans]) + res
			unbr /= ans
		}
		res = string(base[unbr]) + res

		for i := 0; i < len(res); i++ {
			z01.PrintRune(rune(res[i]))
		}
	} else {
		ans := len(base)
		for nbr >= ans {
			res = string(base[nbr%ans]) + res
			nbr /= ans
		}
		res = string(base[nbr]) + res

		for i := 0; i < len(res); i++ {
			z01.PrintRune(rune(res[i]))
		}
	}
}
