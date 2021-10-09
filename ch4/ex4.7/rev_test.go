package rev

import "testing"

func TestRev(t *testing.T) {
    tests := []struct{
        s string
        expect string
    }{
        {"", ""},
        {"hello, 世界!", "!界世 ,olleh"},
    }

    for _, test := range tests {
        r := reverse([]byte(test.s))
        if string(r) != test.expect {
            t.Errorf("rev(%q)=%q, expecting %q\n", test.s, r, test.expect)
        }
    }
}