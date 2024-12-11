package main

import "strings"

func day8Part1() int {
	fileContents, err := ReadFileLines("./data/day8.txt")

	if err != nil {
		return 0
	}

	return navigateWasteland(fileContents)
}

func day8Part2() int {
	fileContents, err := ReadFileLines("./data/day8.txt")

	if err != nil {
		return 0
	}

	return navigateGhostWasteland(fileContents)
}

func navigateWasteland(lines []string) int {
	paths := parsePaths(lines[2:])
	return navigatePaths(paths, lines[0])
}

func navigateGhostWasteland(lines []string) int {
	paths := parsePaths(lines[2:])
	return ghostNavigatePaths(paths, lines[0])

}

func parsePaths(lines []string) map[string][]string {
	paths := make(map[string][]string)

	for _, line := range lines {
		cleanLine := strings.ReplaceAll(line, " = (", ",")
		cleanLine = strings.ReplaceAll(cleanLine, " ", "")
		cleanLine = strings.ReplaceAll(cleanLine, ")", "")
		breakLines := strings.Split(cleanLine, ",")
		paths[breakLines[0]] = []string{breakLines[1], breakLines[2]}
	}

	return paths
}

func navigatePaths(paths map[string][]string, instructions string) int {
	moves := 0
	moveIndex := 0
	current := "AAA"

	for current != "ZZZ" {
		nextMove := rune(instructions[moveIndex])
		if nextMove == 'L' {
			current = paths[current][0]
		} else {
			current = paths[current][1]
		}
		moves++
		moveIndex = (moveIndex + 1) % len(instructions)
	}

	return moves
}

func ghostNavigatePaths(paths map[string][]string, instructions string) int {
	nodes := getGhostStarts(paths)
	movesPerNode := make([]int, len(nodes))

	for i, currentNode := range nodes {
		current := currentNode
		moveIndex := 0
		moves := 0
		for rune(current[2]) != 'Z' {

			nextMove := rune(instructions[moveIndex])
			if nextMove == 'L' {
				current = paths[current][0]
			} else {
				current = paths[current][1]
			}
			moves++
			moveIndex = (moveIndex + 1) % len(instructions)
		}

		movesPerNode[i] = moves
	}
	return LCM(movesPerNode[0], movesPerNode[1], movesPerNode[2:]...)
}

func getGhostStarts(paths map[string][]string) []string {
	starts := []string{}
	for key := range paths {
		if rune(key[2]) == 'A' {
			starts = append(starts, key)
		}
	}

	return starts
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
