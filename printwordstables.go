package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, v := range a {
		for i := 0; i < len(v); i++ {
			z01.PrintRune(rune(v[i]))
		}
		z01.PrintRune('\n')
	}
}
