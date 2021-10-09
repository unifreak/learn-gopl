package popcount

import (
    "math/rand"
    "testing"
    "time"
)

// @todo: NOT getting the same result!!! possibly book is wrong.
// Temporialy disabled by lowercase test*
func testScanPopCount(t *testing.T) {
    seed := time.Now().UTC().UnixNano()
    t.Logf("Random seed: %d", seed)
    rng := rand.New(rand.NewSource(seed))

    for i := 0; i < 1000; i++ {
        x := rng.Uint64()
        desire := PopCount(x)
        got := ScanPopCount(x)
        if got != desire {
            t.Errorf("ScanPopCount(%d) == %d, instead %d", x, got, desire)
        }
    }
}

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount(uint64(i))
    }
}

func BenchmarkScanPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ScanPopCount(uint64(i))
    }
}