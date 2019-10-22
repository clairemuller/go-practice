package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// check that num is greater than zero and not a decimal
func wholeNumCheck(num float64) bool {
	if num < 0 || num != math.Trunc(num) {
		return false
	}
	return true
}

// find the highest power of num that is less than num
func findPower(num float64, base float64) float64 {
	var power float64
	for num >= math.Pow(base, power) {
		power += 1
	}
	return (power - 1)
}

func FormatBinary(num float64) string {
	if !wholeNumCheck(num) {
		return "Whole numbers only!"
	}
	if num == 0 {
		return "0"
	}
	var binarySlice []string
	firstLoop := true
	var power = findPower(num, 2)

	for num >= 0 && power >= 0 {
		// if on the first loop or num is greater/equal to 2^power
		// add 1 to slice and subtract 2^power from num
		// otherwise add a 0
		// decrement power each time
		if firstLoop || num >= math.Pow(2.0, power) {
			binarySlice = append(binarySlice, "1")
			num -= math.Pow(2.0, power)
			firstLoop = false
		} else {
			binarySlice = append(binarySlice, "0")
		}
		power--
	}
	// convert slice of strings into single string
	return strings.Join(binarySlice, "")
}

func FormatHex(num float64) string {
	if !wholeNumCheck(num) {
		return "Whole numbers only!"
	}
	if num == 0 {
		return "0"
	}

	var hexSlice []string
	var power = findPower(num, 16)

	for num >= 0 && power >= 0 {
		// as long as num is greater/equal to 16^power
		// decrement num by 16^power
		// use count to keep track of how many times we decrement
		count := 0
		for num >= math.Pow(16.0, power) {
			num -= math.Pow(16.0, power)
			count++
		}
		// use count to determine what digit/char to add to slice
		if count < 10 {
			hexSlice = append(hexSlice, strconv.Itoa(count))
		} else if count == 10 {
			hexSlice = append(hexSlice, "A")
		} else if count == 11 {
			hexSlice = append(hexSlice, "B")
		} else if count == 12 {
			hexSlice = append(hexSlice, "C")
		} else if count == 13 {
			hexSlice = append(hexSlice, "D")
		} else if count == 14 {
			hexSlice = append(hexSlice, "E")
		} else if count == 15 {
			hexSlice = append(hexSlice, "F")
		}
		power--
	}

	return strings.Join(hexSlice, "")
}

func main() {
	var num float64 = 10000
	// fmt.Println(FormatBinary(num))
	fmt.Println(FormatHex(num))

	// fmt.Printf("%v converted to binary is %v\n", num, FormatBinary(num))
	// fmt.Printf("%v converted to hexadecimal is %v\n", num, FormatHex(num))
}
