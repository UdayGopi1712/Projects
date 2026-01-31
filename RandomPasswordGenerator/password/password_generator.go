package password

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Characters used for password  generation.
var charset = []rune(
	"abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"!@#$%^&*()-_=+[]{}|\\:;\"'<>,.?/",
)

func Generate(length int) string {

	var b strings.Builder
	b.Grow(length)

	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b.WriteRune(charset[n.Int64()])
	}

	return b.String()
}
