package main

import (
	"strconv"
	"strings"
)

func dayTwoPartOne() int {
	fileContents, err := ReadFileLines("./data/day2.txt")

	if err != nil {
		return 0
	}

	return loopThroughGames(fileContents, false)
}

func dayTwoPartTwo() int {
	fileContents, err := ReadFileLines("./data/day2.txt")

	if err != nil {
		return 0
	}

	return loopThroughGames(fileContents, true)
}

func loopThroughGames(lines []string, part2 bool) int {
	total := 0
	for i, val := range lines {
		cleanedGame := removePrefix(val)
		if part2 {
			total += powerOfGame(cleanedGame)
		} else {
			if validateHands(cleanedGame) {
				total += i + 1
			}
		}
	}

	return total
}

func validateHands(line string) bool {
	hands := strings.Split(line, "; ")

	for _, hand := range hands {
		colors := strings.Split(hand, ", ")
		for _, color := range colors {
			if !validateHand(color) {
				return false
			}
		}

	}
	return true
}

func validateHand(hand string) bool {
	vals := strings.Split(hand, " ")
	count, _ := strconv.Atoi(vals[0])
	return isValidColorCounts(vals[1], count)
}

func removePrefix(line string) string {
	idx := strings.Index(line, ":")
	if idx == -1 {
		return line
	}
	return line[idx+2:]
}

func isValidColorCounts(color string, number int) bool {
	if color == "red" {
		return number <= 12
	} else if color == "green" {
		return number <= 13
	} else if color == "blue" {
		return number <= 14
	}
	return false
}

func powerOfGame(line string) int {
	cleanLine := strings.ReplaceAll(line, ";", ",")
	colors := strings.Split(cleanLine, ", ")
	colorCount := map[string]int{
		"red":   1,
		"green": 1,
		"blue":  1,
	}

	for _, score := range colors {
		vals := strings.Split(score, " ")
		color := vals[1]
		count, _ := strconv.Atoi(vals[0])

		if count > colorCount[color] {
			colorCount[color] = count
		}
	}

	return colorCount["red"] * colorCount["blue"] * colorCount["green"]
}
