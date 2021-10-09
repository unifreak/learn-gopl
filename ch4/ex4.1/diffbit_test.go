package diffbit

import "testing"

func TestDiffbit(t *testing.T) {
    tests := []struct{
        a, b [32]byte
        want int
    } {
        {[32]byte{0xff}, [32]byte{0xff}, 0},
        {[32]byte{0x00}, [32]byte{0x00}, 0},
        {[32]byte{0x01}, [32]byte{0x00, 0x01}, 2},
        // 2f: 00101111 f2:11110010
        {[32]byte{0x2f, 0xf2}, [32]byte{0xf2, 0x2f}, 12},
    }
    for _, test := range tests {
        got := DiffBit(test.a, test.b)
        if got != test.want {
            t.Errorf("diff between %b, %b got %d, want %d\n", test.a, test.b, got, test.want)
        }
    }
}