package password

import "testing"

func TestGenerateLength(t *testing.T) {

	pwd := Generate(16)

	if len(pwd) != 16 {
		t.Fatalf("expected length 16, got %d", len(pwd))
	}
}

func TestGenerateRandomness(t *testing.T) {

	p1 := Generate(16)
	p2 := Generate(16)

	if p1 == p2 {
		t.Fatal("passwords should not be equal")
	}
}
