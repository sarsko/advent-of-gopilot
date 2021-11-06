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
	// part 1
	fmt.Println(part1(input))
	// part 2
	fmt.Println(part2(input))
}

// part1 function
func part1(input []string) int {
	var sum int
	for _, line := range input {
		// call check function
		sum += check(line)
	}
	return sum
}

// part2 function
func part2(input []string) int {
	var sum int
	for _, line := range input {
		// call check function
		sum += check2(line)
	}
	return sum
}

// isPalindrome function
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// check2 function
func check2(line string) int {
	// make bool
	var isPal bool
	// keep track of palindromes
	var palindromes []string
	// create inhypernet bool
	var inHypernet bool
	// iterate over line
	for i := 0; i < len(line)-1; i++ {
		// check if letter is "["
		if line[i] == '[' {
			// set inHypernet to true
			inHypernet = true
		}
		// check if letter is "]"
		if line[i] == ']' {
			// set inHypernet to false
			inHypernet = false
		}
		// if not in hypernet
		if !inHypernet {
			// check if i+2 in bounds
			if i+2 < len(line) {
				// check if next three letters are a palindrome
				if isPalindrome(line[i : i+3]) {
					// check if all three letters are the same
					if line[i] == line[i+1] && line[i] == line[i+2] {
						// pass
					} else {
						// set isPal to true
						isPal = true
						// add to palindromes
						palindromes = append(palindromes, line[i:i+3])
					}
				}
			}
		}
	}
	// if not isPal, return 0
	if !isPal {
		return 0
	}
	// set isPal to false
	isPal = false

	// iterate over line
	for i := 0; i < len(line)-1; i++ {
		// check if letter is "["
		if line[i] == '[' {
			// set inHypernet to true
			inHypernet = true
		}
		// check if letter is "]"
		if line[i] == ']' {
			// set inHypernet to false
			inHypernet = false
		}
		// if in hypernet
		if inHypernet {
			// check if next three letters are a palindrome
			if isPalindrome(line[i : i+3]) {
				// check if all three letters are the same
				if line[i] == line[i+1] && line[i] == line[i+2] {
					// pass
				} else {
					// iterate over palindromes
					for _, pal := range palindromes {
						// check if pal[0] == line[i+1] and pal[1] == line[i] and pal[2] == line[i+1]
						if pal[0] == line[i+1] && pal[1] == line[i] && pal[2] == line[i+1] {
							// set isPal to true
							isPal = true
						}
					}

				}
			}
		}
	}
	// if isPal, return 1
	if isPal {
		return 1
	}
	// return 0
	return 0

}

// check function
func check(line string) int {
	// make bool
	var isPal bool
	// iterate over line
	for i := 0; i < len(line)-1; i++ {
		// check if i+3 is in bounds // hack to get i+4 <= len(line)
		if i+3 < len(line) {
			// check if next four letters are a palindrome
			if isPalindrome(line[i : i+4]) {
				// check if all four letters are the same
				if line[i] == line[i+1] && line[i] == line[i+2] && line[i] == line[i+3] {
					// pass
				} else {
					// set isPal to true
					isPal = true
				}
			}
		}
	}
	// if not isPal, return 0
	if !isPal {
		return 0
	}
	// create inhypernet bool
	var inHypernet bool
	// iterate over line
	for i := 0; i < len(line)-1; i++ {
		// check if letter is "["
		if line[i] == '[' {
			// set inHypernet to true
			inHypernet = true
		}
		// check if letter is "]"
		if line[i] == ']' {
			// set inHypernet to false
			inHypernet = false
		}
		// if in hypernet
		if inHypernet {
			// check if next four letters are a palindrome
			if isPalindrome(line[i : i+4]) {
				// check if all four letters are the same
				if line[i] == line[i+1] && line[i] == line[i+2] && line[i] == line[i+3] {
					// pass
				} else {
					// set isPal to false
					isPal = false
				}
			}
		}
	}
	// if isPal, return 1
	if isPal {
		return 1
	}
	// return 0
	return 0

}
