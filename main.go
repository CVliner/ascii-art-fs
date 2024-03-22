package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Check if the number of command-line arguments is sufficient
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . something standard")
		return
	}

	// Extract the input string and banner template from command-line arguments
	input1 := os.Args[1]
	input := ""
	for i := 0; i < len(input1); i++ {
		input += string(input1[i])
	}
	template := os.Args[2]

	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	// Construct the path to the banner file based on the style
	bannerPath := filepath.Join(currentDir, "banners", template+".txt")

	// Check if the file exists
	if _, err := os.Stat(bannerPath); os.IsNotExist(err) {
		log.Println("File does not exist:", err)
		return
	}

	// Read the banner data
	bannerData := ReadBanner(bannerPath)
	if len(bannerData) == 0 {
		fmt.Println("Failed to read banner data.")
		return
	}

	var output string
	// Process each line of the input
	input = strings.ReplaceAll(input, `\n`, "\n")
	lines := strings.Split(input, "\n")
	// fmt.Println(len(input))
	for _, line := range lines {
		if line == "" {
			output += "\n"
			continue
		}
		// Process each character in the line
		for i := 0; i < 8; i++ {
			for _, char := range []byte(line) {
				if int(char) < 32 || int(char) > 126 {
					fmt.Println("Error")
					return
				}
				// Add the ASCII representation of the character to the output
				output += bannerData[1+int(char-32)*9+i]
			}
			// Add a newline character after processing each character
			output += "\n"
		}
	}
	fmt.Print(output)
}

func ReadBanner(path string) []string {
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil
	}

	return strings.Split(string(raw), "\n")
}
