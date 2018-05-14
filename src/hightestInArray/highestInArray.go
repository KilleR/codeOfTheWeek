package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Starting...")

	input := []int{20, 500, 900}
	fmt.Printf("Longest of %v is: %d", input, FindLongest(input))
}

func FindLongest(in []int) (longest int) {
	longestLength := 0

	for i:=0;i< len(in); i++ { // must use FOR because `range` on slice does not always return same order
		vString := strconv.Itoa(in[i])
		kLength := len(vString)
		if kLength > longestLength {
			longestLength = kLength
			longest = in[i]
		}
	}

	return
}

