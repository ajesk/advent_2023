package main

import (
	"strconv"
	"strings"
)

func day9Part1() int {
	fileContents, err := ReadFileLines("./data/day9.txt")
	print(fileContents)

	if err != nil {
		return 0
	}

	return extrapolateValues(fileContents, false)
}

func day9Part2() int {
	fileContents, err := ReadFileLines("./data/day9.txt")

	if err != nil {
		return 0
	}

	return extrapolateValues(fileContents, true)
}

func extrapolateValues(lines []string, backwards bool) int {
	total := 0

	for _, line := range lines {
		total += findNextValueOfLine(line, backwards)
	}

	return total
}

func findNextValueOfLine(line string, backwards bool) int {
	baseNums := convertStringToIntArray(line)
	breakdown := reduceToZeroes(baseNums)
	if backwards {
		return scaleToFindNextValueBackwards(breakdown)
	} else {
		return scaleToFindNextValue(breakdown)
	}
}

func convertStringToIntArray(line string) []int {
	var baseNums = []int{}
	split := strings.Split(line, " ")
	for _, numStr := range split {
		num, _ := strconv.Atoi(numStr)
		baseNums = append(baseNums, num)
	}
	return baseNums
}

func reduceToZeroes(nums []int) [][]int {
	breakdown := [][]int{}
	current := nums
	next := []int{}

	for {
		for i := 0; i < len(current)-1; i++ {
			diff := current[i+1] - current[i]
			next = append(next, diff)
		}

		breakdown = append(breakdown, current)
		current = next
		next = []int{}
		if isCurrentAllZeroes(current) {
			break
		}
	}
	breakdown = append(breakdown, current)

	return breakdown
}

func isCurrentAllZeroes(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}

	return true
}

func scaleToFindNextValue(breakdown [][]int) int {
	total := 0
	for _, level := range breakdown {
		total += level[len(level)-1]
	}

	return total
}

func scaleToFindNextValueBackwards(breakdown [][]int) int {
	leftMost := 0
	for i := len(breakdown) - 2; i >= 0; i-- {
		leftMost = breakdown[i][0] - leftMost
	}

	return leftMost
}
