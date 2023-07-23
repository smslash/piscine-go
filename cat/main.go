package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		for _, r := range string(data) {
			z01.PrintRune(r)
		}
	} else {
		for _, arg := range args {
			data, err := ioutil.ReadFile(arg)
			if err != nil {
				z01.PrintRune('E')
				z01.PrintRune('R')
				z01.PrintRune('R')
				z01.PrintRune('O')
				z01.PrintRune('R')
				z01.PrintRune(':')
				z01.PrintRune(' ')
				for _, r := range err.Error() {
					z01.PrintRune(r)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			for _, r := range string(data) {
				z01.PrintRune(r)
			}
		}
	}
}
