package popcount_test

import (
	"testing"

	popcount "./"
)

func BenchmarkSingleExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x123456789ABCDEF)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(0x123456789ABCDEF)
	}
}
