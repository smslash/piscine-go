package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 1 {
		return
	}
	ans := false
	if len(args) > 0 && args[0] == "--upper" {
		ans = true
		args = os.Args[2:]
	}

	for _, arg := range args {
		n := 0
		space := 0
		for _, c := range arg {
			if c < '0' || c > '9' {
				z01.PrintRune(' ')
				space++
				break
			}
			n = n*10 + int(c-'0')
		}
		if n < 1 || n > 26 {
			if space == 0 {
				z01.PrintRune(' ')
			}
			continue
		}
		if ans {
			z01.PrintRune(rune(n + 64))
		} else {
			z01.PrintRune(rune(n + 96))
		}
	}
	z01.PrintRune('\n')
}
