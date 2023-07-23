package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for j := len(args) - 1; j >= 0; j-- {
		for i := 0; i < len(args[j]); i++ {
			z01.PrintRune(rune(args[j][i]))
		}
		z01.PrintRune('\n')
	}
}
