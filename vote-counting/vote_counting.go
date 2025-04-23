package votecounting

import (
	"fmt"
	"os"
	"strings"
)

// Processes voting from a text file
func fromFileName(fname string) string {
	if !strings.HasSuffix(fname, ".votes") {
		fmt.Println("Error reading file from filename; filename has no file extension .votes")
		return ""
	}
	filepath := fname
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file from filename; file not found.")
		return ""
	}

	if len(data) == 0 {
		fmt.Println("Error reading file from filename; file is empty")
	}

	return string(data)
}

// Process voting from user input
func fromUserInput() string {
	return "" // TODOs
}
