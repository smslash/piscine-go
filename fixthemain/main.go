package main

import "github.com/01-edu/z01"

const (
	OPEN  = false
	CLOSE = true
)

type Door struct {
	state bool
}

func OpenDoor(ptrDoor *Door) {
	s := ("Door Opening...\n")
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) bool {
	s := ("Door Closing...\n")
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	s := ("is the Door opened ?\n")
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	s := ("is the Door closed ?\n")
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
