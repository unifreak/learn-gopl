package rotate

import "testing"

func TestRotate(t *testing.T) {
    tests := []struct{
        s []int
        pos int
        want []int
    }{
        {[]int{0, 1, 2}, 1, []int{2, 0, 1}},
        {[]int{0, 1, 2}, 2, []int{1, 2, 0}},
        {[]int{0, 1, 2, 3, 4, 5}, 2, []int{4, 5, 0, 1, 2, 3}},
        {[]int{0, 1, 2, 3, 4, 5}, 8, []int{4, 5, 0, 1, 2, 3}},
        {[]int{0, 1, 2, 3, 4, 5}, 0, []int{0, 1, 2, 3, 4, 5}},
        {[]int{0, 1, 2, 3, 4, 5}, -2, []int{2, 3, 4, 5, 0, 1}},
        {[]int{0, 1, 2, 3, 4, 5}, -8, []int{2, 3, 4, 5, 0, 1}},
    }
    for _, test := range tests {
        s := test.s
        r := rotate(test.s, test.pos)
        if !equal(r, test.want) {
            t.Errorf("rotate(%v, %d)=%v, want %v", s, test.pos, r, test.want)
        }
    }
}

func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, _ := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}