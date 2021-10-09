package rev

import "testing"

func TestRev(t *testing.T) {
    tests := []struct{
        in [N]int
        out [N]int
    }{
        {[N]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, [N]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
        {[N]int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}, [N]int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0}},
        {[N]int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, [N]int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}},
        {[N]int{10, 0, 0, 0, 0, 0, 0, 0, 0, 0}, [N]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 10}},
    }
    for _, test := range tests {
        in := test.in
        reverse(&test.in)
        if test.in != test.out {
            t.Errorf("reverse(%v)=%v, expection %v\n", in, test.in, test.out)
        }
    }
}