package main

import (
	"strconv"
	"log"
	"fmt"
)

func main() {
	fmt.Println(ReverseNumber(-123.456))
}

func ReverseNumber(origin float64) (reversed float64) {
	// preserve negativity
	isNegative := origin < 0

	// store the positive version of the number
	posNumber := origin
	if isNegative {
		posNumber = posNumber * -1
	}

	// stringify the value - no exponent, all digits, 64bit
	strNumber := strconv.FormatFloat(posNumber, 'f', -1, 64)

	// flip the string
	var revStrNumber string
	for _, v := range strNumber {
		revStrNumber = string(v) + revStrNumber // add the next rune to the front
	}

	// convert back to a number
	revPosNumber, err := strconv.ParseFloat(revStrNumber, 64)
	if err != nil {
		log.Fatalln("Failed to parse reversed number", err)
	}

	// re-apply the negative if necessary - to the output value `reversed`
	reversed = revPosNumber
	if isNegative {
		reversed = revPosNumber * -1
	}

	return
}
