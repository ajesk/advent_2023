package main

import (
	"strconv"
	"strings"
	"unicode"
)

func dayOnePartOne() int {
	fileContents, err := ReadFileLines("./data/one/reals.txt")

	if err != nil {
		return 0
	}

	return getSumOfAllLines(fileContents)
}

func getSumOfAllLines(lines []string) int {
	total := 0
	for _, val := range lines {
		total += getNumberOfLine(val)
	}

	return total
}

func getNumberOfLine(line string) int {
	stripped := removeAllNonNumbers(line)
	firstChar := stripped[0]
	lastChar := stripped[len(stripped)-1]

	str, _ := strconv.Atoi(string(firstChar) + string(lastChar))
	return str
}

func removeAllNonNumbers(line string) string {
	var result strings.Builder

	for _, char := range line {
		if unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func dayOnePartTwo() int {
	fileContents, err := ReadFileLines("./data/one/real.txt")

	if err != nil {
		return 0
	}

	return sanitizeAndGetSumOfAllLines(fileContents)
}

func sanitizeAndGetSumOfAllLines(lines []string) int {
	total := 0
	for _, line := range lines {
		cleanLine := replaceFirstNumber(line)
		cleanLineBack := replaceLastNumber(line)
		firstLineVal := getFirstNumberOfLine(cleanLine)
		lastLineVal := getLastNumberOfLine(cleanLineBack)
		val, _ := strconv.Atoi(firstLineVal + lastLineVal)
		total += val
	}

	return total
}

func getFirstNumberOfLine(line string) string {
	stripped := removeAllNonNumbers(line)
	firstChar := stripped[0]
	return string(firstChar)
}

func getLastNumberOfLine(line string) string {
	stripped := removeAllNonNumbers(line)
	lastChar := stripped[len(stripped)-1]
	return string(lastChar)
}

func replaceFirstNumber(line string) string {
	strToInt := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	firstIdx := 999
	var first string

	for key := range strToInt {
		idx := strings.Index(line, key)
		if idx != -1 && idx < firstIdx {
			first = key
			firstIdx = idx
		}
	}

	result := strings.Replace(line, first, strToInt[first], 1)
	return result
}

func replaceLastNumber(line string) string {
	strToInt := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	lastIdx := -1
	var last string

	for key := range strToInt {
		idxEnd := strings.LastIndex(line, key)
		if idxEnd != -1 && idxEnd > lastIdx {
			last = key
			lastIdx = idxEnd
		}
	}

	if lastIdx == -1 {
		return line
	}

	return line[:lastIdx] + strToInt[last] + line[lastIdx+len(last):]
}
