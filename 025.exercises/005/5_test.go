package main

import "testing"

func TestSqrt(t *testing.T) {
	f := 1.2
	if f < 0 {
		t.Errorf("Expact f(%v) > 0", f)
	}
}
