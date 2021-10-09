package main

import "testing"


func TestUnits(t *testing.T) {
    if MB / KB != 1000 {
        t.Error("MB / KB != 3", MB, KB)
    }
    if ZB / EB != 1000 {
        t.Error("ZB / EB != 3, = %v", ZB, EB, ZB/EB) // constant overflow int
    }
}