// @? undefined: Popcount when run
//      go test ch2/ex2.3/popcount_test.go
// @a cd to dir and run
//      go test -bench .
// @? why this work but not providing path?
package popcount

import (
    "testing"
)

const maxUint64 = 1<<64 - 1

var funcs = map[string]func(uint64)int{
    "PopCount": PopCount,
    "LoopPopCount": LoopPopCount,
    "ScanPopCount": ScanPopCount,
    "ShiftPopCount": ShiftPopCount,
}

func TestPopCountFuncs(t *testing.T) {
    tc := []struct{
        in uint64
        want int
    }{
        {0, 0},
        {1, 1},
        {7, 3},
        {64, 1},
        {uint64(maxUint64), 64},
    }

    for _, c := range tc {
        for fname, fn := range funcs {
            if out := fn(c.in); out != c.want {
                t.Errorf("%s(%v)=%v want %v\n", fname, c.in, out, c.want)
            }
        }
    }
}

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount(uint64(i))
    }
}

func BenchmarkLoopPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        LoopPopCount(uint64(i))
    }
}

func BenchmarkShiftPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ShiftPopCount(uint64(i))
    }
}

func BenchmarkScanPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ScanPopCount(uint64(i))
    }
}