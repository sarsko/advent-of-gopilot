package main

import (
	"io/ioutil"
	"strings"
)

func read_input() string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(input)
}

// main function
func main() {
	input := read_input()
	// Part 1
	part1(input)
	// Part 2
	part2(input)
}

// Part 1
func part1(input string) {

	// split the input into an array of strings
	input_array := strings.Split(input, "\n")
	// iterate over the array
	for _, value := range input_array {
		// current position starts at 5
		position := 5
		// iterate over each character
		for _, char := range string(value) {
			// if char is U
			if char == 'U' {
				// if position is not 1, 2 or 3
				if position != 1 && position != 2 && position != 3 {
					// decrease position by 3
					position -= 3
				}
			}
			// if char is D
			if char == 'D' {
				// if position is not 7, 8 or 9
				if position != 7 && position != 8 && position != 9 {
					// increase position by 3
					position += 3
				}
			}
			// if char is L
			if char == 'L' {
				// if position is not 1, 4, 7
				if position != 1 && position != 4 && position != 7 {
					// decrease position by 1
					position -= 1

				}
			}
			//	if char is R
			if char == 'R' {
				// if position is not 3, 6, 9
				if position != 3 && position != 6 && position != 9 {
					// increase position by 1
					position += 1
				}
			}
		}
		// print position
		println(position)
	}
	// print newline
	println()
}

// Part 2
func part2(input string) {
	// split the input into an array of strings
	input_array := strings.Split(input, "\n")
	// iterate over the array
	for _, value := range input_array {
		// current position starts at 5
		position := 5
		// iterate over each character
		for _, char := range string(value) {
			// if char is U
			if char == 'U' {
				// if position is not 5, 2, 1, 4, 9
				if position != 5 && position != 2 && position != 1 && position != 4 && position != 9 {
					// decrease position by 4
					position -= 4
				}
				// if position is -1
				if position == -1 {
					// set position to 1
					position = 1
				}
				// if position is 9
				if position == 9 {
					// set position to 11
					position = 11
				}
			}
			// if char is D
			if char == 'D' {
				// if position is not 5, 10, 13, 12, 9
				if position != 5 && position != 10 && position != 13 && position != 12 && position != 9 {
					// increase position by 4
					position += 4
				}
				// if position is 15
				if position == 15 {
					// set position to 13
					position = 13
				}
				// if position is 5
				if position == 5 {
					// set position to 3
					position = 3
				}
			}
			// if char is L
			if char == 'L' {
				// if position is not 5, 2, 10, 13, 1
				if position != 5 && position != 2 && position != 10 && position != 13 && position != 1 {
					// decrease position by 1
					position -= 1
				}
			}
			//	if char is R
			if char == 'R' {
				// if position is not 1, 4, 9, 12, 13
				if position != 1 && position != 4 && position != 9 && position != 12 && position != 13 {
					// increase position by 1
					position += 1
				}
			}
		}
		// print position
		println(position)
	}
	// print newline
	println()
}
