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

// ScanPopCount calculate pop count by scan bit from right to left.
func ScanPopCount(x uint64) int {
    r := 0
    for x != 0 {
        x = x&(x-1)
        r++
    }
    return r
}