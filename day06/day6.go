package main

import (
	"bufio"
	"fmt"
	"os"
)

// read input from input.txt. return a slice of strings
func readInput() []string {
	var input []string
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

// main function
func main() {
	input := readInput()
	part1(input)
}

//part 1
func part1(input []string) {
	// make an empty output string
	var output string
	// make another empty string
	var output2 string
	// iterate with j for len 6
	for j := 0; j < len(input[0]); j++ {
		// make array of length 26 init to 0
		var count [26]int
		// iterate over input with indexes
		for i := 0; i < len(input); i++ {
			// get char at index j
			char := input[i][j]
			// increment count at index of char
			count[char-'a']++
		}
		// find max index of count
		max := 0
		for i := 0; i < 26; i++ {
			if count[i] > count[max] {
				max = i
			}
		}
		// find min index of count
		min := 0
		for i := 0; i < 26; i++ {
			if count[i] < count[min] {
				min = i
			}
		}
		// append min char to output2
		output2 += string(min + 'a')
		// append max to output
		output += string(max + 'a')
	}
	// print output
	fmt.Println(output)
	// print output2
	fmt.Println(output2)
}
