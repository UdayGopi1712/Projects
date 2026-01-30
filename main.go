package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var char = map[int]rune{
		// lowercase a-z
		0: 'a', 1: 'b', 2: 'c', 3: 'd', 4: 'e',
		5: 'f', 6: 'g', 7: 'h', 8: 'i', 9: 'j',
		10: 'k', 11: 'l', 12: 'm', 13: 'n', 14: 'o',
		15: 'p', 16: 'q', 17: 'r', 18: 's', 19: 't',
		20: 'u', 21: 'v', 22: 'w', 23: 'x', 24: 'y',
		25: 'z',

		// uppercase A-Z
		26: 'A', 27: 'B', 28: 'C', 29: 'D', 30: 'E',
		31: 'F', 32: 'G', 33: 'H', 34: 'I', 35: 'J',
		36: 'K', 37: 'L', 38: 'M', 39: 'N', 40: 'O',
		41: 'P', 42: 'Q', 43: 'R', 44: 'S', 45: 'T',
		46: 'U', 47: 'V', 48: 'W', 49: 'X', 50: 'Y',
		51: 'Z',

		// special characters (password-safe)
		52: '!', 53: '@', 54: '#', 55: '$', 56: '%',
		57: '^', 58: '&', 59: '*', 60: '(', 61: ')',
		62: '-', 63: '_', 64: '=', 65: '+',
		66: '[', 67: ']', 68: '{', 69: '}',
		70: '|', 71: '\\',
		72: ':', 73: ';',
		74: '"', 75: '\'',
		76: '<', 77: '>', 78: '?',
		79: '/', 80: '.', 81: ',',
	}

	fmt.Printf("Enter the length of the password:")
	// by default generates 16 digit length
	var length_password int = 16
	fmt.Scan(&length_password)

	fmt.Printf("Enter number of random passwords to be created: ")
	//default value is 1
	var no_of_passwords int = 1
	fmt.Scan(&no_of_passwords)

	// to avoid same passwords
	rand.Seed(time.Now().UnixNano())

	// main logic will start here
	for i := 0; i < no_of_passwords; i++ {
		var password string = ""
		for j := 0; j < length_password; j++ {
			password += string(char[rand.Int()%81])
		}
		fmt.Printf("password %d : %s\n", i+1, password)
	}

}
