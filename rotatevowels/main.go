package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Vowels(s []string) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == 'a' || s[i][j] == 'e' || s[i][j] == 'i' || s[i][j] == 'o' || s[i][j] == 'u' {
				return true
			}
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	s := os.Args[1:]
	vowels := []rune{}
	res := []rune{}
	sum := 0

	if Vowels(s) {
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(s[i]); j++ {
				if s[i][j] == 'a' || s[i][j] == 'e' || s[i][j] == 'i' || s[i][j] == 'o' || s[i][j] == 'u' || s[i][j] == 'A' || s[i][j] == 'E' || s[i][j] == 'I' || s[i][j] == 'O' || s[i][j] == 'U' {
					vowels = append(vowels, rune(s[i][j]))
				}
			}
		}

		for i := 0; i < len(s); i++ {
			for j := 0; j < len(s[i]); j++ {
				if s[i][j] == 'a' || s[i][j] == 'e' || s[i][j] == 'i' || s[i][j] == 'o' || s[i][j] == 'u' || s[i][j] == 'A' || s[i][j] == 'E' || s[i][j] == 'I' || s[i][j] == 'O' || s[i][j] == 'U' {
					sum++
					res = append(res, vowels[len(vowels)-sum])
				} else {
					res = append(res, rune(s[i][j]))
				}
			}
			if i+1 != len(s) {
				res = append(res, ' ')
			}
		}

		for i := 0; i < len(res); i++ {
			z01.PrintRune(rune(res[i]))
		}
	} else {
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(s[i]); j++ {
				z01.PrintRune(rune(s[i][j]))
			}
			if i+1 != len(s) {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
