package piscine

import "github.com/01-edu/z01"

func Number(nb int) {
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

func PrintComb() {
	for a := 1; a < 1000; a++ {
		i := a / 100
		j := a / 10 % 10
		k := a % 10

		if i == 7 && j == 8 && k == 9 {
			z01.PrintRune('7')
			z01.PrintRune('8')
			z01.PrintRune('9')
			z01.PrintRune('\n')
			break
		} else if i < j && j < k {
			Number(i)
			Number(j)
			Number(k)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}
