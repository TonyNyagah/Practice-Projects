package main

import (
	"fmt"
	"strconv"
)

func convert_to_decimal(num int) int {
	var current_number int
	var result int = 0
	var total int
	const base int = 2

	// turn the number into a string
	num_string := strconv.Itoa(num)

	// turn the string into an array for easy looping
	num_array := []rune(num_string)

	// loop through the array
	for i := 0; i < len(num_array); i++ {
		// convert the rune to an int
		current_number, _ = strconv.Atoi(string(num_array[i]))

		total = result*base + current_number

		result = total
	}

	fmt.Println(result)
	return result
}

func main() {
	convert_to_decimal(100111)
}
