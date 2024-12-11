package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileLines(fileName string) ([]string, error) {
	var lines []string

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %s", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file: %s", err)
	}

	return lines, nil
}
