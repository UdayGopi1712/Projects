package password

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func cryptoShuffle(b []byte) error {
	for i := len(b) - 1; i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		j := int(n.Int64())
		b[i], b[j] = b[j], b[i]
	}
	return nil
}

func Generate(length int) (string, error) {
	if length < 8 {
		return "", errors.New("length must be >= 8")
	}

	// Characters used for password  generation.
	const (
		lower   = "abcdefghijklmnopqrstuvwxyz"
		upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits  = "0123456789"
		symbols = "!@#$%^&*()-_=+[]{}|\\:;\"'<>,.?/"
	)

	all := lower + upper + digits + symbols

	b := make([]byte, length)

	// enforce policy
	sets := []string{lower, upper, digits, symbols}
	for i := 0; i < len(sets); i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(sets[i]))))
		if err != nil {
			return "", err
		}
		b[i] = sets[i][n.Int64()]
	}

	// fill remaining
	for i := len(sets); i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(all))))
		if err != nil {
			return "", err
		}
		b[i] = all[n.Int64()]
	}

	// shuffle
	if err := cryptoShuffle(b); err != nil {
		return "", err
	}

	return string(b), nil
}
