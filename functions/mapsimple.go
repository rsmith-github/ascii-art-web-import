package functions

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func MakeMapSimple(input string, w http.ResponseWriter, r *http.Request) {

	fil := r.FormValue("dropdown")

	// Read file using ioutil
	content, _ := os.Open(fil)

	// Use bufio scanner to scan file content line by line.
	scanner := bufio.NewScanner(content)

	// Variable to keep track of position in file.
	Rune := 31
	// Map to store characters.
	ascii_map := make(map[rune][]string)

	// For every line in stabdard.txt
	for scanner.Scan() {
		// If empty string, increment Rune to access next key.
		if scanner.Text() == "" {
			Rune++
			// Else append line by line to the array of strings.
		} else {
			ascii_map[rune(Rune)] = append(ascii_map[rune(Rune)], scanner.Text())
		}
	}

	// Replace literal new line with \\n
	input = strings.ReplaceAll(input, "\r\n", "\\n")

	// Check if space is present.
	backn_checker := 0
	// Array to split by \n later.
	var split []string
	// Edge case.
	if input == "\\n" {
		fmt.Fprintf(w, "\n")
		return
		// If there are any \n in input, and the last character is not \n
	} else if strings.Contains(input, "\\n") || strings.Contains(input, "\n") /* && string(input[len(input)-2]) != "\\" && string(input[len(input)-1]) != "n" */ {
		backn_checker++
		// Split by \n.
		if strings.Contains(input, "\\n") {
			split = strings.Split(input, "\\n")
		}
		// Loop over split array.
		for j := 0; j < len(split); j++ {
			// Loop over each line in ascii character.
			for i := 0; i < 8; i++ {
				// Print one line of every character the input.
				for _, char := range split[j] {
					fmt.Fprintf(w, "%s", ascii_map[char][i])
				}
				// Print new line to go to next row.
				fmt.Fprintf(w, "\n")
			}
		}
		return
	} /* else if input == "" {

	} */

	// Edge case variable. If it is equal to 1, print a new line.
	edge := 0
	// If the last two characters are \ and n, remove them from the input variable and increment edge case.
	if string(input[len(input)-2]) == "\\" && string(input[len(input)-1]) == "n" {
		input = input[:len(input)-2]
		edge++
	}

	// if there is no \n, print the ascii art normally.
	if backn_checker != 1 {
		// Print ascii character.
		for i := 0; i < 8; i++ {
			for _, char := range input {
				fmt.Fprintf(w, "%s", ascii_map[char][i])
			}
			fmt.Fprintf(w, "\n")
		}
	}
	// If edge case is met, print new line.
	if edge == 1 {
		fmt.Fprintf(w, "\n")
	}
	return
}
