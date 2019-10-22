package main

import (
	"fmt"
	"math"
	"strconv"
)

func findPower(num float64, base float64) float64 {
	var power float64
	for num >= math.Pow(base, power) {
		power += 1
	}
	return (power - 1)
}

func FormatBinary(num float64) string {
	// check that num is greater than zero and not a decimal
	if num < 0 || num != math.Trunc(num) {
		return "Whole numbers only!"
	}
	if num == 0 {
		return "0"
	}
	var binarySlice []int
	firstLoop := true
	var power = findPower(num, 2)

	for num >= 0 && power >= 0 {
		// if on the first loop or num is greater/equal to 2^power
		// add 1 to slice and subtract 2^power from num
		// otherwise add a 0
		// decrement power each time
		if firstLoop || num >= math.Pow(2.0, power) {
			binarySlice = append(binarySlice, 1)
			num -= math.Pow(2.0, power)
			firstLoop = false
		} else {
			binarySlice = append(binarySlice, 0)
		}
		power--
	}
	// convert slice of numbers into a string
	var binaryString string
	for i := 0; i < len(binarySlice); i++ {
		binaryString += strconv.Itoa(binarySlice[i])
	}
	return binaryString
}

func toHex(num int) int {
	return num + 100
}

func main() {
	var num float64 = 9
	fmt.Println(FormatBinary(num))

	// fmt.Printf("%v converted to binary is %v\n", num, FormatBinary(num))
	// fmt.Printf("%v converted to hexadecimal is %v\n", num, toHex(num))
}
