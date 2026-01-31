package password

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

var charset = []rune(
	"abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"!@#$%^&*()-_=+[]{}|\\:;\"'<>,.?/",
)

func Generate(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("password length must be > 0")
	}

	var b strings.Builder
	b.Grow(length)

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b.WriteRune(charset[n.Int64()])
	}

	return b.String(), nil
}
