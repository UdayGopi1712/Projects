package password

import (
	"testing"
	"unicode"
)

func TestGenerateLength(t *testing.T) {

	pwd, err := Generate(16)

	if err != nil {
		t.Fatalf("expected no error but got error: %s", err)
	}

	if len(pwd) != 16 {
		t.Fatalf("expected length 16, got %d", len(pwd))
	}
}

func TestGenerateRandomness(t *testing.T) {

	p1, err1 := Generate(16)
	p2, err2 := Generate(16)

	if err1 != nil {
		t.Fatalf("expected no error but got error: %s", err1)
	}

	if err2 != nil {
		t.Fatalf("expected no error but got error: %s", err2)
	}

	if p1 == p2 {
		t.Fatal("passwords should not be equal")
	}
}

func TestGeneratePolicy(t *testing.T) {
	pwd, err := Generate(16)
	if err != nil {
		t.Fatal(err)
	}

	hasLower, hasUpper, hasDigit, hasSymbol := false, false, false, false

	for _, c := range pwd {
		switch {
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsDigit(c):
			hasDigit = true
		default:
			hasSymbol = true
		}

		if hasDigit && hasLower && hasSymbol && hasUpper {
			break
		}
	}

	if !(hasLower && hasUpper && hasDigit && hasSymbol) {
		t.Fatal("password does not meet complexity requirements")
	}
}
