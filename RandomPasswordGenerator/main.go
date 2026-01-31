package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"test/password"
)

func main() {

	// default values for length of password and no.of passwords if not provided.
	length := flag.Int("length", 16, "Password length (min 8)")
	count := flag.Int("count", 1, "Number of passwords")
	flag.Parse()

	if *length < 8 {
		log.Fatal("length must be >= 8")
	}

	if *count < 1 {
		log.Fatal("count must be >= 1")
	}

	for i := 0; i < *count; i++ {

		pwd, err := password.Generate(*length)
		if err != nil {
			fmt.Printf("Error %s", err)
			os.Exit(1)
		}

		fmt.Println(pwd)
	}
}
