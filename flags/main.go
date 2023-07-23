package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func PrintTerminal() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Printf("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func Insert(s []string) {
	after := []rune{}
	before := []rune{}
	for i := 0; i < len(s); i++ {
		if s[i][0] == '-' && s[i][1] == '-' && s[i][2] == 'i' && s[i][3] == 'n' && s[i][4] == 's' && s[i][5] == 'e' && s[i][6] == 'r' && s[i][7] == 't' {
			for j := 9; j < len(s[i]); j++ {
				after = append(after, rune(s[i][j]))
			}
		} else if s[i][0] == '-' && s[i][1] == 'i' {
			for j := 3; j < len(s[i]); j++ {
				after = append(after, rune(s[i][j]))
			}
		} else if s[i][0] == '-' && s[i][1] == '-' && s[i][2] == 'o' && s[i][3] == 'r' && s[i][4] == 'd' && s[i][5] == 'e' && s[i][6] == 'r' {
			continue
		} else if s[i][0] == '-' && (s[i][1] == 'o' || s[i][1] == 'h') {
			continue
		} else if s[i][0] == '-' && s[i][1] == '-' && s[i][2] == 'h' && s[i][3] == 'e' && s[i][4] == 'l' && s[i][5] == 'p' {
			continue
		} else if len(s[i]) < 1 {
			continue
		} else {
			for j := 0; j < len(s[i]); j++ {
				before = append(before, rune(s[i][j]))
			}
		}
	}

	res := []rune{}

	for i := 0; i < len(before); i++ {
		res = append(res, before[i])
	}

	for i := 0; i < len(after); i++ {
		res = append(res, after[i])
	}

	for i := 0; i < len(res); i++ {
		z01.PrintRune(res[i])
	}
}

func InOrder(s []string) {
	after := []rune{}
	before := []rune{}
	for i := 0; i < len(s); i++ {
		if s[i][0] == '-' && s[i][1] == '-' && s[i][2] == 'i' && s[i][3] == 'n' && s[i][4] == 's' && s[i][5] == 'e' && s[i][6] == 'r' && s[i][7] == 't' {
			for j := 9; j < len(s[i]); j++ {
				after = append(after, rune(s[i][j]))
			}
		} else if s[i][0] == '-' && s[i][1] == 'i' {
			for j := 3; j < len(s[i]); j++ {
				after = append(after, rune(s[i][j]))
			}
		} else if s[i][0] == '-' && s[i][1] == '-' && s[i][2] == 'o' && s[i][3] == 'r' && s[i][4] == 'd' && s[i][5] == 'e' && s[i][6] == 'r' {
			continue
		} else if s[i][0] == '-' && (s[i][1] == 'o' || s[i][1] == 'h') {
			continue
		} else if s[i][0] == '-' && s[i][1] == '-' && s[i][2] == 'h' && s[i][3] == 'e' && s[i][4] == 'l' && s[i][5] == 'p' {
			continue
		} else if len(s[i]) < 1 {
			continue
		} else {
			for j := 0; j < len(s[i]); j++ {
				before = append(before, rune(s[i][j]))
			}
		}
	}

	res := []rune{}

	for i := 0; i < len(before); i++ {
		res = append(res, before[i])
	}

	for i := 0; i < len(after); i++ {
		res = append(res, after[i])
	}

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}

	for i := 0; i < len(res); i++ {
		z01.PrintRune(res[i])
	}
}

func main() {
	var s []string
	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i]) >= 1 {
			s = append(s, os.Args[i])
		}
	}

	count_help := 0
	count_insert := 0
	count_order := 0

	for _, v := range s {
		if v == "--help" || v == "-h" {
			count_help++
		}

		if len(v) > 7 {
			if v[0] == '-' && v[1] == '-' && v[2] == 'i' && v[3] == 'n' && v[4] == 's' && v[5] == 'e' && v[6] == 'r' && v[7] == 't' {
				count_insert++
			}
		}

		if len(v) > 1 {
			if v[0] == '-' && v[1] == 'i' {
				count_insert++
			}
		}

		if v == "--order" || v == "-o" {
			count_order++
		}
	}

	if (count_help == 1 || count_help == 0) && count_insert == 0 && count_order == 0 {
		PrintTerminal()
	}

	if count_help == 0 && count_insert == 1 && count_order == 0 {
		Insert(s)
	}

	if count_help == 0 && count_insert == 0 && count_order == 1 {
		InOrder(s)
	}

	if count_help == 0 && count_insert == 1 && count_order == 1 {
		InOrder(s)
	}

	z01.PrintRune('\n')
}
