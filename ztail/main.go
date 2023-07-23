package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	result := 0
	index := args[2]
	status := 0

	for i := 0; i < len(index); i++ {
		result *= 10
		result += int(index[i] - 48)
	}

	if len(args) == 4 {
		file, err := os.ReadFile(args[3])
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			status = 1
		} else {
			fmt.Printf(string(file[len(file)-result:]))
		}

	} else {
		for i := 3; i < len(args); i++ {
			file, err := os.ReadFile(args[i])
			if err != nil {
				fmt.Printf("%v\n", err.Error())
				status = 1
			} else {
				if i != 3 {
					fmt.Printf("\n")
				}
				fmt.Printf("==> %v <==\n", args[i])

				if len(string(file)) <= result {
					fmt.Printf("%v", string(file))
					status = 1
				} else {
					fmt.Printf(string(file[len(file)-result:]))
				}

			}

		}
	}
	os.Exit(status)
}
