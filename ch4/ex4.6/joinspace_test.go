package joinspace

import "testing"

func TestJoinspace(t *testing.T) {
    tests := []struct{
        s string
        expect string
    }{
        {"", ""},
        {" ", " "},
        {" \t", " "},
        {`你
           好 `, "你 好 "},
    }

    for _, test := range tests {
        r := joinspace([]byte(test.s))
        if (string(r) != test.expect) {
            t.Errorf("joinspace(%q)=%q, expecting %q\n", test.s, r, test.expect)
        }
    }
}