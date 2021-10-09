package main

import "testing"
import "os"
import "bytes"

func TestWordfreq(t *testing.T) {
    tests := []struct{
        args []string
        want string
    }{
        {
            []string{"wordfreq", "data.txt"},
            "Hello!: 1\nagain: 2\nagain...: 1\nand: 2\n世界: 1\n"},
    }
    for _, test := range tests {
        os.Args = test.args
        out = new(bytes.Buffer)
        // @?: how to factor out and mock scanner?
        main()
        got := out.(*bytes.Buffer).String()
        if got != test.want {
            t.Errorf("%q = %q, want %q", os.Args, got, test.want)
        }
    }
}