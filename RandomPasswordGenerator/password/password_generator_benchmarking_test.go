package password

import "testing"

func BenchmarkGenerate16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Generate(16)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerate64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Generate(64)
		if err != nil {
			b.Fatal(err)
		}
	}
}
