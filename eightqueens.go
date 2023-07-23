package piscine

import "github.com/01-edu/z01"

func Check(board [8]int, x int) {
	if x == 8 {
		for _, v := range board {
			z01.PrintRune(rune(v + 1 + 48))
		}
		z01.PrintRune(rune(10))
		return
	}

	for y := 0; y < 8; y++ {
		ans := true
		for i := 0; i < x; i++ {
			if board[i] == y {
				ans = false
			}
			if board[i]+i == y+x {
				ans = false
			}
			if board[i]-i == y-x {
				ans = false
			}
			if !ans {
				break
			}
		}
		if ans {
			board[x] = y
			Check(board, x+1)
		}
	}
}

func EightQueens() {
	board := [8]int{}
	Check(board, 0)
}
