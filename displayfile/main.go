package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	s := os.Args[1]
	data, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf(string(data))
}
