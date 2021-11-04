package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Read the file 'input.txt'
func read_input() string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(input)
}

// the main function
func main() {
	// init a counter with value 0
	var counter int = 0
	// read the input
	input := read_input()
	// split the input into lines
	lines := strings.Split(input, "\n")
	// for line in the lines
	for _, line := range lines {
		// split the line into words
		words := strings.Split(line, " ")
		// array of ints
		var ints []int
		// for each word in the words
		for _, word := range words {
			// if the word is ""
			if word == "" {
				// skip it
				continue
			}
			// convert the word to int
			i, err := strconv.Atoi(word)
			if err != nil {
				panic(err)
			}
			// add the int to the ints array
			ints = append(ints, i)
		}
		// sort the ints array
		sort(ints)
		// if ints[0] + ints[1] > ints[2]
		if ints[0]+ints[1] > ints[2] {
			// increase	the counter
			counter++
		}
	}
	// print the counter
	println(counter)
	// set the counter to 0
	counter = 0
	// make three int arrays
	var a, b, c []int
	// for line in the lines
	for _, line := range lines {
		// split the line into words
		words := strings.Split(line, " ")
		// array_counter
		var array_counter int = 0

		// for each word in the words
		for _, word := range words {
			// if the word is ""
			if word == "" {
				// skip it
				continue
			}
			// convert the word to int
			i, err := strconv.Atoi(word)
			if err != nil {
				panic(err)
			}
			// if array_counter is 0
			if array_counter == 0 {
				// add the int to the a array
				a = append(a, i)
			}
			// if array_counter is 1
			if array_counter == 1 {
				// add the int to the b array
				b = append(b, i)
			}
			// if array_counter is 2
			if array_counter == 2 {
				// add the int to the c array
				c = append(c, i)
			}
			// increase the array_counter
			array_counter++
		}
		// if the a array has 3 ints
		if len(a) == 3 {
			// sort the a array
			sort(a)
			// if a[0] + a[1] > a[2]
			if a[0]+a[1] > a[2] {
				// increase the counter
				counter++
			}
			// reset the a array
			a = []int{}
		}
		// if the b array has 3 ints
		if len(b) == 3 {
			// sort the b array
			sort(b)
			// if b[0] + b[1] > b[2]
			if b[0]+b[1] > b[2] {
				// increase the counter
				counter++
			}
			// reset the b array
			b = []int{}
		}
		// if the c array has 3 ints
		if len(c) == 3 {
			// sort the c array
			sort(c)
			// if c[0] + c[1] > c[2]
			if c[0]+c[1] > c[2] {
				// increase the counter
				counter++
			}
			// reset the c array
			c = []int{}
		}
	}
	// print the counter
	println(counter)
}

// sort the ints array
func sort(ints []int) {
	// for i in the ints array
	for i := 0; i < len(ints)-1; i++ {
		// if the int at i is greater than the int at i+1
		if ints[i] > ints[i+1] {
			// swap the ints at i and i+1
			ints[i], ints[i+1] = ints[i+1], ints[i]
		}
	}
}
