package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = 223372036854775808
		} else {
			n = -n
		}
	}
	if n/10 != 0 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10) + '0')
}
