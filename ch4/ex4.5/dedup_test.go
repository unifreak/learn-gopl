package dedup

import "testing"

func TestDedup(t *testing.T) {
    tests := []struct{
        s []string
        want []string
    }{
        {[]string{}, []string{}},
        {[]string{"a", "a"}, []string{"a"}},
        {[]string{"hi我", "hi你", "hi你"}, []string{"hi我", "hi你"}},
        {[]string{"aba", "aba", "c", "c", "aba", "aba", "aba"}, []string{"aba", "c", "aba"}},
    }

    for _, test := range tests {
        s := test.s
        r := dedup(test.s)
        if !equal(r, test.want) {
            t.Errorf("dedup(%v)=%v, want %v\n", s, r, test.want)
        }
    }
}

func equal(a, b []string) bool {
    if (len(a) != len(b)) {
        return false
    }
    for i, _ := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}