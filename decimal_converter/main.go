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

func toBinary(num float64) int {
	if num == 0 {
		return 0
	}
	var binaryNum int
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
	// then convert string to an int
	var binaryString string
	for i := 0; i < len(binarySlice); i++ {
		binaryString += strconv.Itoa(binarySlice[i])
	}
	binaryNum, _ = strconv.Atoi(binaryString)
	return binaryNum
}

func toHex(num int) int {
	return num + 100
}

func main() {
	var num float64 = 8
	fmt.Println(toBinary(num))

	// fmt.Printf("%v converted to binary is %v\n", num, toBinary(num))
	// fmt.Printf("%v converted to hexadecimal is %v\n", num, toHex(num))
}
