package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read the file 'input.txt'
func read_input() string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}

	return input
}

func main() {
	// take the result of read_input()
	input := read_input()
	// split the string into an array of strings
	// split on ", "
	input_array := strings.Split(input, ", ")
	// keep track of the current direction as int
	direction := 0
	// and the current x and y position as int
	x := 0
	y := 0

	// keep track of places we've been
	visited := make(map[string]bool)

	// for each string in input_array
	for _, value := range input_array {
		// if the string starts with "R"
		if strings.HasPrefix(value, "R") {
			// increase the direction by 1
			direction++
			// if the direction is greater than 3
			if direction > 3 {
				// set the direction to 0
				direction = 0
			}
		} else if strings.HasPrefix(value, "L") {
			// decrease the direction by 1
			direction--
			// if the direction is less than 0
			if direction < 0 {
				// set the direction to 3
				direction = 3
			}
		}
		// take the rest of the string and convert it to an int
		distance, _ := strconv.Atoi(value[1:])
		// if the direction is 0
		if direction == 0 {
			// increase the x position by the int iteratively
			for i := 0; i < distance; i++ {
				// add the current position to the visited map
				visited[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
				x++
				// if we've been here before
				if visited[strconv.Itoa(x)+","+strconv.Itoa(y)] {
					// print the distance from the start
					fmt.Println(abs(x) + abs(y))
				}

			}
		}
		// if the direction is 1
		if direction == 1 {
			// increase the y position by the int iteratively
			for i := 0; i < distance; i++ {
				// add the current position to the visited map
				visited[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
				y++
				// if we've been here before
				if visited[strconv.Itoa(x)+","+strconv.Itoa(y)] {
					// print the distance from the start
					fmt.Println(abs(x) + abs(y))
				}

			}
		}
		// if the direction is 2
		if direction == 2 {
			// decrease the x position by the int iteratively
			for i := 0; i < distance; i++ {
				// add the current position to the visited map
				visited[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
				x--
				// if we've been here before
				if visited[strconv.Itoa(x)+","+strconv.Itoa(y)] {
					// print the distance from the start
					fmt.Println(abs(x) + abs(y))
					// add the current position to the visited map
					visited[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
				}

			}
		}
		// if the direction is 3
		if direction == 3 {
			// decrease the y position by the int iteratively
			for i := 0; i < distance; i++ {
				// add the current position to the visited map
				visited[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
				y--
				// if we've been here before
				if visited[strconv.Itoa(x)+","+strconv.Itoa(y)] {
					// print the distance from the start
					fmt.Println(abs(x) + abs(y))
				}

			}
		}

	}
	// print the manhattan distance
	fmt.Println(abs(x) + abs(y))
}

// make a function abs which calculate the absolute value of an int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
