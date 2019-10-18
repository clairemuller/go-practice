package main

import (
	"fmt"
	"math"
)

func toBinary(num float64) float64 {
	// find highest power of 2 that is less than num
	// start with 2 power of 0
	// var power float64 = 0

	// while num is less than or equal to 2^power, increase power
	for num >= math.Pow(2, power) {
		// fmt.Printf("2^%v = %v\n", power, math.Pow(2, power))
		power += 1
	}


	// var spaces

	return power
}

func toHex(num float64) float64 {
	return num + 100
}

func main() {
	var num float64 = 34
	fmt.Println(toBinary(num))

	// fmt.Printf("%v converted to binary is %v\n", num, toBinary(num))
	// fmt.Printf("%v converted to hexadecimal is %v\n", num, toHex(num))
}
