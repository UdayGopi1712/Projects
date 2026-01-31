package main

import (
	"fmt"
	"os"
	"strconv"

	"test/password"
)

func main() {
	length := 16
	count := 1

	if len(os.Args) >= 2 {
		if v, err := strconv.Atoi(os.Args[1]); err == nil {
			length = v
		}
	}
	if len(os.Args) >= 3 {
		if v, err := strconv.Atoi(os.Args[2]); err == nil {
			count = v
		}
	}

	for i := 0; i < count; i++ {
		pwd, err := password.Generate(length)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(pwd)
	}
}
