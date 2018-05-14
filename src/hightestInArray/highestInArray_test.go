package main

import "testing"

func TestFindLongest(t *testing.T) {
	var(
		input []int
		l int
	)
	input = []int{20, 500, 900} // test case for later higher, but not longer
	l = FindLongest(input)
	if l != 500 {
		t.Fatalf("Failure, result should be %d, instead got %d", 500, l)
	}

	input = []int{200, 300, 1000} // test case for later higher, and longer
	l = FindLongest(input)
	if l != 1000 {
		t.Fatalf("Failure, result should be %d, instead got %d", 1000, l)
	}
}
