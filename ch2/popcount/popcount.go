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
	return int(pc[byte(x>>0*8)] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// for Ex2.3
func LoopPopCount(x uint64) int {
	var n int
	var i uint
	for i = 0; i < 8; i++ {
		n += int(pc[byte(x>>i*8)])
	}
	return n
}

// for Ex2.4
func ShiftPopCount(x uint64) int {
	var n int
	var i uint
	for i = 0; i < 64; i++ {
		n += int(x & 1)
		x >>= 1
	}
	return n
}

// for Ex2.5 @todo