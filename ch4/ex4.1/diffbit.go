// Write a function that counts the number of bits that are different in two SH256
// hashes. (See PopCount)
package diffbit

import "crypto/sha256"

var pc [256]byte

// Init precompute a lookup table. See popcount
func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func DiffBit(a, b [sha256.Size]byte) int {
    n := 0
    for i := 0; i < sha256.Size; i++ {
        n += int(pc[a[i]^b[i]])
    }
    return n
}