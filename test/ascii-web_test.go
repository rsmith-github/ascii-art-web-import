package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

func ReadTestFile() []string {
	data, err := ioutil.ReadFile("test-output.txt")
	if err != nil {
		panic(err)
	}

	contentString := string(data)

	contentSplit := strings.Split(contentString, "#")

	return contentSplit
}

func TestAsciiArt(t *testing.T) {
	testCases := []string{
		"{123}\r\n<Hello> (World)!",
		"123??",
		"$% \"=",
		"123 T/fs#R",
	}

	expectedOutput := ReadTestFile()

	for index, testCase := range testCases {
		cmd := exec.Command("go", "run", ".", testCase)

		output, _ := cmd.Output()

		fmt.Println("Expected:")
		fmt.Println(expectedOutput[index])
		fmt.Println("Found:")
		fmt.Println(string(output))

		if string(output) != expectedOutput[index] {
			t.Errorf("\nTest fails when given case:\n\t\"%s\","+"\nThe test should show:\n%s\nInstead it shows:\n%s", testCase, expectedOutput[index], output)
		}
	}
}
