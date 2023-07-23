package piscine

import "github.com/01-edu/z01"

func Number2(nb int) {
	if nb == 0 {
		z01.PrintRune('0')
	} else if nb == 1 {
		z01.PrintRune('1')
	} else if nb == 2 {
		z01.PrintRune('2')
	} else if nb == 3 {
		z01.PrintRune('3')
	} else if nb == 4 {
		z01.PrintRune('4')
	} else if nb == 5 {
		z01.PrintRune('5')
	} else if nb == 6 {
		z01.PrintRune('6')
	} else if nb == 7 {
		z01.PrintRune('7')
	} else if nb == 8 {
		z01.PrintRune('8')
	} else if nb == 9 {
		z01.PrintRune('9')
	}
}

func PrintComb2() {
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			if a < b {
				i := a / 10
				j := a % 10
				x := b / 10
				y := b % 10
				if i == 9 && j == 8 && x == 9 && y == 9 {
					z01.PrintRune('9')
					z01.PrintRune('8')
					z01.PrintRune(' ')
					z01.PrintRune('9')
					z01.PrintRune('9')
					z01.PrintRune('\n')
					break
				}
				Number2(i)
				Number2(j)
				z01.PrintRune(' ')
				Number2(x)
				Number2(y)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
