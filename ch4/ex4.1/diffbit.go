package diffbit

var pc [256]byte

// Init precompute a lookup table. See popcount
func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func DiffBit(a, b [32]byte) int {
    n := 0
    for i := 0; i < 32; i++ {
        if a[i] == b[i] {
            continue
        }
        n += int(pc[a[i]^b[i]])
    }
    return n
}