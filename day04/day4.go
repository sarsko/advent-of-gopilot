package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// read input.txt
func readInput() []string {
	var input []string
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

// main
func main() {
	input := readInput()
	// result strings
	var result []string
	// result ints
	var resultInts []int
	// get result and resultints from getcorrectrooms
	result, resultInts = getCorrectRooms(input)
	// call part2 on the result of getCorrectRooms
	part2(result, resultInts)
}

// split lines
func splitLines(input []string) []string {
	var lines []string
	for _, line := range input {
		lines = append(lines, line)
	}
	return lines
}

//count occurences of char in string
func countOccurences(char rune, line string) int {
	count := 0
	for _, c := range line {
		if c == char {
			count++
		}
	}
	return count
}

// remove duplicates from string
func removeDuplicates(s string) string {
	var result string
	for _, c := range s {
		if !strings.ContainsRune(result, c) {
			result += string(c)
		}
	}
	return result
}

// sort ints
func sortInts(ints []int) []int {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i] > ints[j] {
				ints[i], ints[j] = ints[j], ints[i]
			}
		}
	}
	return ints
}

//get correct rooms. returns a tuple of strings and numbers
func getCorrectRooms(input []string) ([]string, []int) {
	// array for old real rooms
	var oldRooms []string
	// array for numbers
	var numbers []int
	// init a counter with value 0
	var counter int = 0
	// split input into lines
	lines := splitLines(input)
	// for each line
	for _, line := range lines {
		// get 6 last characters of line
		last6 := line[len(line)-6:]
		// remove last character of last6
		last6 = last6[:len(last6)-1]
		// remove last 7 characters of line
		line = line[:len(line)-7]
		// get last 3 characters of line
		last3 := line[len(line)-3:]
		// convert last3 to int
		last3Int, _ := strconv.Atoi(last3)
		// remove last 3 from line
		line = line[:len(line)-3]
		// keep old line
		oldLine := line
		// remove last from oldLine
		oldLine = oldLine[:len(oldLine)-1]
		// strip line of "-"
		line = strings.Replace(line, "-", "", -1)

		// create an array of ints
		var ints []int
		// for each character in last6
		for _, char := range last6 {
			// count occurences of char in line
			count := countOccurences(char, line)
			// append count to ints
			ints = append(ints, count)
		}
		// remove duplicates from line. assign to new variable
		newLine := removeDuplicates(line)
		// create an array of ints
		var ints2 []int
		// for each character in newLine
		for _, char := range newLine {
			// count occurences of char in line
			count := countOccurences(char, line)
			// append count to ints2
			ints2 = append(ints2, count)
		}
		// sort ints2
		ints2 = sortInts(ints2)
		// reverse ints2
		for i, j := 0, len(ints2)-1; i < j; i, j = i+1, j-1 {
			ints2[i], ints2[j] = ints2[j], ints2[i]
		}
		// ok bool false
		ok := false
		// check if first 5 elements of ints2 are equal to first 5 elements of ints
		if ints[0] == ints2[0] && ints[1] == ints2[1] && ints[2] == ints2[2] && ints[3] == ints2[3] && ints[4] == ints2[4] {
			// set ok to true
			ok = true
		}

		// init previous to a large number
		previous := 1000000

		// for idx and int in ints
		for idx, i := range ints {
			// if i is larger than previous
			if i > previous {
				// set ok to false
				ok = false
				// break
				break
			}
			// else if i is equal to previous
			if i == previous {
				// check if last6[idx] is less than last6[idx-1]
				if last6[idx] < last6[idx-1] {
					// set ok to false
					ok = false
					// break
					break
				}
			}
			// check if i is smaller than previous
			if i < previous {
				// set previous to i
				previous = i
			}
		}
		// if ok is true
		if ok {
			// add oldLine to oldRooms
			oldRooms = append(oldRooms, oldLine)
			// add last3Int to numbers
			numbers = append(numbers, last3Int)
			// increase counter by last3Int
			counter += last3Int
		}
	}
	// print counter
	println(counter)
	// return oldRooms and numbers
	return oldRooms, numbers
}

// shift string
func shiftString(s string, shift int) string {
	// shift modulo 26
	shift = shift % 26
	// init result
	result := ""
	// for each char in s
	for _, c := range s {
		// get char code
		code := int(c)
		// if code + shift is greater than 122
		if code+shift > 122 {
			// add (code + shift) - 26 to result
			result += string(code + shift - 26)
		} else {
			// add code + shift to result
			result += string(code + shift)
		}
	}
	// return result
	return result
}

//part 2
func part2(oldRooms []string, numbers []int) {
	// iterate over oldRooms. keep a counter
	var counter int = 0
	for _, oldRoom := range oldRooms {
		// shift all characters in oldRoom by number[counter]
		shifted := shiftString(oldRoom, numbers[counter])
		// print shifted
		// if shifted contains north
		if strings.Contains(shifted, "north") {
			// print shifted
			println(shifted)
			// print number[counter]
			println(numbers[counter])
		}
		// increase counter
		counter++
	}

}
