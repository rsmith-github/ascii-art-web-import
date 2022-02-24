package functions

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

func MakeMapSimple(input string, w http.ResponseWriter, r *http.Request) (string, error) {

	// get name of selected banner file from radio button.
	fil := r.FormValue("option")

	// Open file.
	content, err := os.Open(fil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

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

	if input == " " {
		return strings.Join(ascii_map[' '], "\n"), nil
	} else if len(input) == 1 {
		return strings.Join(ascii_map[rune(input[0])], "\n"), nil
	}

	// Replace literal new line with \\n
	input = strings.ReplaceAll(input, "\r\n", "NEWLINEhere")

	result := ``

	// Check if \n is present.
	// Array to split by \n later.
	var split []string

	// If there are any new lines in input
	if strings.Contains(input, "NEWLINEhere") {
		// Split by \n.
		split = strings.Split(input, "NEWLINEhere")
		// Loop over split array.
		for j := 0; j < len(split); j++ {
			// Loop over each line in ascii character.
			for i := 0; i < 8; i++ {
				// Print one line of every character the input.
				for _, char := range split[j] {
					// fmt.Fprintf(w, "%s", ascii_map[char][i])
					result += ascii_map[char][i]
				}
				// add new line to go to next row.
				result += "\n"
			}
		}
		return result, nil
	} else if input == "" {
		// Add error message?
		return "", nil
	}

	// Print ascii character.
	for i := 0; i < 8; i++ {
		for _, char := range input {
			// fmt.Fprintf(w, "%s", ascii_map[char][i])
			result += ascii_map[char][i]
		}
		// fmt.Fprintf(w, "\n")
		result += "\n"
	}

	return result, nil
}
