package main

import (
	"fmt"
	"os"
	"strconv"

	"test/password"
)

func main() {
	// default values for length of password and no.of passwords if not provided.
	length := 16
	count := 1

	// checking for commands line arguments for length and no.of passwords
	if len(os.Args) >= 2 {
		if v, err := strconv.Atoi(os.Args[1]); err == nil {
			if v < 8 {
				fmt.Println("length must be greater than 8 for secure passwords.")
				os.Exit(0)
			}
			length = v
		}
	}
	if len(os.Args) >= 3 {
		if v, err := strconv.Atoi(os.Args[2]); err == nil {
			if v <= 0 {
				fmt.Println("Mininum no.of passwords should be 1 or more")
				os.Exit(0)
			}
			count = v
		}
	}

	for i := 0; i < count; i++ {

		pwd := password.Generate(length)

		fmt.Println(pwd)
	}
}
