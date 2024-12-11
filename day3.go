package main

import (
	"strconv"
	"unicode"
)

func dayThreePartOne() int {
	fileContents, err := ReadFileLines("./data/day3.txt")

	if err != nil {
		return 0
	}

	return countNumbersWithAdjacentSymbols(fileContents)
}

func countNumbersWithAdjacentSymbols(lines []string) int {
	currentNumber := ""
	isCurrentNearSymbol := false
	total := 0
	for y, line := range lines {
		for x, char := range line {
			if unicode.IsDigit(char) {
				currentNumber += string(char)
				isCurrentNearSymbol = isCurrentNearSymbol || scanForSymbol(lines, y, x)
			} else {
				if len(currentNumber) != 0 && isCurrentNearSymbol {
					currentInt, _ := strconv.Atoi(currentNumber)
					total += currentInt
				}

				currentNumber = ""
				isCurrentNearSymbol = false
			}
		}
	}

	if len(currentNumber) != 0 && isCurrentNearSymbol {
		currentInt, _ := strconv.Atoi(currentNumber)
		total += currentInt
	}

	return total
}

func scanForSymbol(lines []string, y int, x int) bool {
	h := len(lines)
	w := len(lines[0])

	moves := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, move := range moves {
		nextY := y + move[0]
		nextX := x + move[1]

		if nextY < 0 || nextY == h || nextX < 0 || nextX == w {
			continue
		}

		if isSymbol(rune(lines[nextY][nextX])) {
			return true
		}
	}

	return false
}

func isSymbol(ch rune) bool {
	return !unicode.IsDigit(ch) && !unicode.IsLetter(ch) && ch != '.'
}

func dayThreePartTwo() int {
	fileContents, err := ReadFileLines("./data/day3.txt")

	if err != nil {
		return 0
	}

	return countNumbersAdjacentToGears(fileContents)
}

func countNumbersAdjacentToGears(lines []string) int {
	total := 0
	for y, line := range lines {
		for x, char := range line {
			if char == '*' {
				total += multiplySurroundingNumbers(lines, y, x)
			}
		}
	}

	return total
}

func multiplySurroundingNumbers(lines []string, y int, x int) int {
	surroundingNumbers := getSurroundingNumbers(lines, y, x)
	if len(surroundingNumbers) == 2 {
		return surroundingNumbers[0] * surroundingNumbers[1]
	}
	return 0
}

func getSurroundingNumbers(lines []string, y int, x int) []int {

	surroundingCells := findAllSurroundingCells(lines, y, x)

	return findNumbersFromSurroundingCells(lines, surroundingCells)

}

func findNumbersFromSurroundingCells(lines []string, cells [][]int) []int {
	h := len(lines)
	w := len(lines[0])
	result := []int{}
	var visited = make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	for _, cell := range cells {
		y := cell[0]
		x := cell[1]
		if visited[y][x] {
			continue
		}
		visited[y][x] = true

		currentNumStr := string(lines[y][x])
		var currentCell rune
		pointer := x - 1
		for pointer >= 0 {
			visited[y][pointer] = true
			currentCell = rune(lines[y][pointer])

			if !unicode.IsDigit(currentCell) {
				break
			}

			currentNumStr = string(currentCell) + currentNumStr
			pointer--
		}

		// go right
		pointer = x + 1
		for pointer < w {
			visited[y][pointer] = true
			currentCell = rune(lines[y][pointer])

			if !unicode.IsDigit(currentCell) {
				break
			}

			currentNumStr += string(currentCell)
			pointer++
		}

		currentInt, _ := strconv.Atoi(currentNumStr)
		result = append(result, currentInt)
	}

	return result
}

func findAllSurroundingCells(lines []string, y int, x int) [][]int {
	var numbers [][]int
	h := len(lines)
	w := len(lines[0])
	moves := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, move := range moves {
		nextY := y + move[0]
		nextX := x + move[1]

		if nextY < 0 || nextY == h || nextX < 0 || nextX == w {
			continue
		}

		if unicode.IsDigit(rune(lines[nextY][nextX])) {
			arr := []int{nextY, nextX}
			numbers = append(numbers, arr)
		}
	}

	return numbers
}
