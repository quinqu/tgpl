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

func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShift(0x123456789ABCDEF)
	}
}
