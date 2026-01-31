package password

import "testing"

func BenchmarkGenerate16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Generate(16)
	}
}

func BenchmarkGenerate64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Generate(64)
	}
}
