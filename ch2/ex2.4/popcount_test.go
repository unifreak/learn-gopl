package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount(uint64(i))
    }
}

func BenchmarkShiftPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ShiftPopCount(uint64(i))
    }
}