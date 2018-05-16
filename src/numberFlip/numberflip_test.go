package main

import "testing"

func TestReverseNumber(t *testing.T) {
	tests := []struct {
		in, expect float64
	}{
		{123, 321}, // simple reverse
		{-123, -321}, // reverse negative
		{123.456, 654.321}, // float reverse
		{-123.456, -654.321}, // negative float reverse
		{1.234, 432.1}, // left-heavy float reverse
		{-1.234, -432.1}, // left-heavy negative float reverse
		{123.4, 4.321}, // right-heavy float reverse
		{-123.4, -4.321}, // right-heavy negative float reverse
	}

	for _, test := range tests {
		output := ReverseNumber(test.in)
		if output != test.expect {
			t.Errorf("Reverse of %g failed: expect %g got %g", test.in, test.expect, output)
		} else {
			t.Logf("Reverse of %g success: expect %g got %g", test.in, test.expect, output)
		}
	}
}