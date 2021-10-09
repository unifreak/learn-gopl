package main

import "testing"

func TestCounttype(t *testing.T) {
    tests := []struct{
        in string
        expect map[string]int
    }{
        {
            "Hello, 世界!!!",
            map[string]int{
                "Letter": 7,
                "Print": 12,
                "Upper": 1,
                "Lower": 4,
                "Punct": 4,
                "Space": 1,
                "Graphic": 12,
            },
        },
    }
    for _, test := range tests {
        for _, r := range test.in {
            countType(r)
        }
        if !equal(test.expect, countsPerType) {
            t.Errorf("countType %q got %v, want %v\n", test.in, countsPerType, test.expect)
        }
    }
}

func equal(a, b map[string]int) bool {
    if (len(a) != len(b)) {
        return false
    }
    for k, av := range a {
        if bv, ok := b[k]; !ok || bv != av {
            return false
        }
    }
    return true
}