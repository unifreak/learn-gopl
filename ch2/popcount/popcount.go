package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the populatioin count (number of set bites) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Ex2.3: Rewrite PopCount to use a loop instead of a signle expression. Compare the
// performance of the two version.
func LoopPopCount(x uint64) int {
	var n int
	for i := 0; i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

// Ex2.4: Write a version of PopCount that counts bits by shifting its argument
// through 64 bit positions, testing the rightmost bit each time. Compare its
// performance to the table-loopup version.
func ShiftPopCount(x uint64) int {
	var n int
	for i := 0; i < 64; i++ {
		n += int(x & 1)
		x >>= 1
	}
	return n
}

// Ex2.5: The expression x&(x-1) clears the rightmost non-zero bit of x. Write
// a version of PopCount that counts bits by using this fact, and assess its performance.
func ScanPopCount(x uint64) int {
    var n int
    for x != 0 {
        x = x&(x-1)
        n++
    }
    return n
}