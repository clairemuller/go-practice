package main

import (
	"fmt"
	"math"
	"strconv"
)

func findPower(num float64, base int) int {
	var power float64
	b := float64(base)
	for num >= math.Pow(b, power) {
		power += 1
	}
	return int(power - 1)
}

func toBinary(num float64) int {
	if num == 0 {
		return 0
	}
	var binaryNum int
	var binarySlice []int
	firstLoop := true
	var position int

	for num >= 0 {
		var power = findPower(num, 2)
		if firstLoop {
			// set the length of the binary number by adding zeroes
			zeroes := power
			binarySlice = append(binarySlice, 1)
			for zeroes > 0 {
				binarySlice = append(binarySlice, 0)
				zeroes--
			}
			firstLoop = false
		} else {
			position = len(binarySlice)
			position -= int(power)
			position--
			binarySlice[position] = 1
		}
		num -= math.Pow(2.0, float64(power))
	}

	// convert binarySlice into binaryNum
	// binaryNum = int(binarySlice)
	var binaryString string
	for i := 0; i < len(binarySlice); i++ {
		binaryString += strconv.Itoa(binarySlice[i])
	}
	fmt.Println(binaryString)
	// fmt.Println(binaryNum)
	return binaryNum
}

func toHex(num int) int {
	return num + 100
}

func main() {
	var num float64 = 19
	fmt.Println(toBinary(num))

	// fmt.Printf("%v converted to binary is %v\n", num, toBinary(num))
	// fmt.Printf("%v converted to hexadecimal is %v\n", num, toHex(num))
}
