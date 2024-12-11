package main

import (
	"strconv"
	"strings"
)

func day12Part1() int {
	fileContents, err := ReadFileLines("./data/day12.txt")
	print(fileContents)

	if err != nil {
		return 0
	}

	return 0
}

func day12Part2() int {
	fileContents, err := ReadFileLines("./data/day12.txt")
	print(fileContents)

	if err != nil {
		return 0
	}

	return 0
}

func splitData(lines []string) ([][]int, [][]int) {
	springConditionLines := [][]int{}
	springCounts := [][]int{}

	for _, line := range lines {
		springRow := strings.Split(line, " ")
		springConditionLines = append(springConditionLines, breakDownSpringCounts(springRow[0]))

		brokenCounts := []int{}
		for _, springConditionNum := range strings.Split(springRow[1], "") {
			count, _ := strconv.Atoi(springConditionNum)
			brokenCounts = append(brokenCounts, count)
		}
	}

	return springConditionLines, springCounts
}

func breakDownSpringCounts(springString string) []int {
	counts := []int{}
	currentCount := 0
	for _, ch := range springString {
		if ch == '?' {
			currentCount++
		} else if currentCount > 0 {
			counts = append(counts, currentCount)
			currentCount = 0
		}
	}

	return counts
}

func dp(springConditions, brokenCounts []int) int {
	if len(brokenCounts) == 0 {
		return 1
	}

	return 0
}
