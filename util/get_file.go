package util

import (
	"bufio"
	"fmt"
	"os"
)

// Get votes from a text file
func FromFileName(fname string) []string {
	filepath := fname
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file from filename; file not found\n", err)
		return nil
	}
	defer file.Close()

	var lines []string
	sc := bufio.NewScanner(file)
	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file from filename\n", err)
	}

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if len(lines) == 0 {
		fmt.Println("Error reading file from filename; file is empty")
		return nil
	}

	return lines
}

// Get votes from user input
func FromUserInput() string {
	return "" // TODOs
}
