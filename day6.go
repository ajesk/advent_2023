package main

import (
	"strconv"
	"strings"
)

func day6Part1() int {
	fileContents, err := ReadFileLines("./data/day6.txt")

	if err != nil {
		return 0
	}

	return compileRecordRanges(fileContents)
}

func compileRecordRanges(lines []string) int {
	times := parseLineNumbers(lines[0])
	records := parseLineNumbers(lines[1])
	total := 1

	for i := range times {
		total *= findRecordRange(times[i], records[i])
	}

	return total
}

func findRecordRange(timeLimit int, record int) int {
	upper := findUpperLimit(timeLimit, record)
	lower := findLowerLimit(timeLimit, record)
	return upper - lower + 1
}

func findUpperLimit(timeLimit int, record int) int {
	for i := timeLimit; i >= 0; i-- {
		result := i * (timeLimit - i)
		if result > record {
			return i
		}
	}
	return -1
}

func findLowerLimit(timeLimit int, record int) int {
	for i := 0; i < timeLimit; i++ {
		result := i * (timeLimit - i)
		if result > record {
			return i
		}
	}
	return -1
}

func parseLineNumbers(line string) []int {
	strs := strings.Split(line, " ")[1:]
	nums := []int{}

	for _, item := range strs {
		if item == "" {
			continue
		}
		num, _ := strconv.Atoi(strings.Trim(item, " "))
		nums = append(nums, num)
	}

	return nums
}

func day6Part2() int {
	fileContents, err := ReadFileLines("./data/day6.txt")

	if err != nil {
		return 0
	}

	return compileLargeRecord(fileContents)
}

func compileLargeRecord(lines []string) int {
	time := squishNumber(lines[0])
	record := squishNumber(lines[1])
	total := findRecordRange(time, record)

	return total
}

func squishNumber(str string) int {
	stripped := strings.ReplaceAll(str[strings.Index(str, ":")+1:], " ", "")
	val, _ := strconv.Atoi(stripped)
	return val
}
