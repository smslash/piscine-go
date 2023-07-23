package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1:]

	for _, v := range s {
		if v == "galaxy" || v == "01" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
