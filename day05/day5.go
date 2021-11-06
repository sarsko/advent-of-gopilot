package main

import (
	"crypto/md5"
	"fmt"
)

// main is the entry point for the application. // LOOOL
func main() {
	puzzleInput := []byte("reyedfim") // manually entered
	// make a counter
	counter := 0
	// build up a string
	password := ""
	// build up another string
	password2 := ""
	// add " " to the end of the string eight times
	for i := 0; i < 8; i++ {
		// add " " to the end of password2
		password2 += " "
	} // Lol
	// while not done
	for {
		// make a hash
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", puzzleInput, counter)))
		// take the hash and format it with fmt.Sprintf. store it
		// in a string
		hashString := fmt.Sprintf("%x", hash)
		// if the first 5 characters of the hash are zeros
		if hashString[0:5] == "00000" {
			// store the 6th character of the hash in a string
			char := string(hashString[5])
			// add the hash of index 6 to the password
			password += string(hashString[5])
			// if password is 8 characters long
			if len(password) == 8 {
				// print the password
				fmt.Println(password)
			}
			// print the character
			fmt.Println(char)
			// check if character is a number
			if char >= "0" && char <= "7" {
				// convert the character to an int
				index := int(char[0] - '0')
				// check if password2 has a space at the index
				if password2[index] == ' ' {
					// if it does, add the 6th character of the hash
					// to the password2
					password2 = password2[:index] + string(hashString[6]) + password2[index+1:]
				}
				// print the password2
				fmt.Println(password2)
			}
		}
		// increment the counter
		counter++
	}

}
