package main

import (
	"fmt"
	"reflect"
)

func day10Part1() int {
	fileContents, err := ReadFileLines("./data/day10.txt")
	print(fileContents)

	if err != nil {
		return 0
	}

	res, _ := bfs(fileContents)

	return res
}

func day10Part2() int {
	return day10Part2File("./data/day10.txt")
}

func day10Part2File(file string) int {
	fileContents, err := ReadFileLines(file)
	print(fileContents)

	if err != nil {
		return 0
	}

	_, double := bfs(fileContents)
	return countInnerSpacesInDouble(double)
}

var moves = map[rune][][]int{
	'S': {{1, 0}, {-1, 0}, {0, 1}, {0, -1}},
	'-': {{0, -1}, {0, 1}},
	'|': {{-1, 0}, {1, 0}},
	'L': {{-1, 0}, {0, 1}},
	'J': {{-1, 0}, {0, -1}},
	'7': {{0, -1}, {1, 0}},
	'F': {{0, 1}, {1, 0}},
}

func bfs(pipes []string) (int, [][]rune) {
	q := [][]int{}
	double := createDoubleArray(pipes)
	count := 1
	var visited = make([][]bool, len(pipes))
	for i := range pipes {
		var rowVisit = make([]bool, len(pipes[0]))
		visited[i] = rowVisit
	}
	start := findStartS(pipes)
	visited[start[0]][start[1]] = true
	q = findDirectionsToStart(pipes, start[0], start[1], q, double)
	q = enqueue(q, []int{})

	for len(q) != 0 {
		current := q[0]
		q = dequeue(q)

		if len(current) == 0 {
			if len(q) == 0 {
				break
			}
			count++
			q = enqueue(q, []int{})
			continue
		}

		y := current[0]
		x := current[1]
		visited[y][x] = true
		currentPipe := rune(pipes[y][x])

		for _, move := range moves[currentPipe] {
			nextY := move[0] + y
			nextX := move[1] + x
			if nextY >= 0 && nextY < len(pipes) &&
				nextX >= 0 && nextX < len(pipes[0]) &&
				!visited[nextY][nextX] {
				moveInDoubleArray(double, move, x*2, y*2)
				q = enqueue(q, []int{nextY, nextX})
			}
		}
	}

	return count, double
}

func findDirectionsToStart(pipes []string, y int, x int, q [][]int, double [][]rune) [][]int {
	for _, move := range moves['S'] {
		nextY := move[0] + y
		nextX := move[1] + x
		if nextY >= 0 && nextY < len(pipes) && nextX >= 0 && nextX < len(pipes[0]) {
			neighbor := rune(pipes[nextY][nextX])
			if reflect.DeepEqual(move, []int{0, 1}) {
				if neighbor != 'J' && neighbor != '7' && neighbor != '-' {
					continue
				}
			} else if reflect.DeepEqual(move, []int{1, 0}) {
				if neighbor != 'J' && neighbor != 'L' && neighbor != '|' {
					continue
				}
			} else if reflect.DeepEqual(move, []int{0, -1}) {
				if neighbor != 'F' && neighbor != 'L' && neighbor != '-' {
					continue
				}
			} else if reflect.DeepEqual(move, []int{-1, 0}) {
				if neighbor != 'F' && neighbor != '7' && neighbor != '|' {
					continue
				}
			}
			moveInDoubleArray(double, move, x*2, y*2)
			q = enqueue(q, []int{nextY, nextX})
		}
	}

	return q
}

func findStartS(pipes []string) []int {
	for y := 0; y < len(pipes); y++ {
		for x := 0; x < len(pipes[0]); x++ {
			ch := pipes[y][x]
			if ch == 'S' {
				return []int{y, x}
			}
		}
	}

	return []int{-1, -1}
}

func moveInDoubleArray(double [][]rune, move []int, x, y int) [][]rune {
	nextX := move[1] + x
	nextY := move[0] + y
	double[nextY][nextX] = 'X'
	nextX += move[1]
	nextY += move[0]
	double[nextY][nextX] = 'X'

	return double
}

func createDoubleArray(pipes []string) [][]rune {
	double := [][]rune{}
	for _, line := range pipes {
		row := []rune{}
		for _, ch := range line {
			row = append(row, ch)
			row = append(row, '0')
		}
		double = append(double, row)

		row = []rune{}
		for range line {
			row = append(row, '0')
			row = append(row, '0')
		}
		double = append(double, row)
	}

	return double
}

func printDouble(double [][]rune) {
	for _, line := range double {
		for _, ch := range line {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}

func countInnerSpacesInDouble(double [][]rune) int {
	var visited = make([][]bool, len(double))
	for i := range double {
		var rowVisit = make([]bool, len(double[0]))
		visited[i] = rowVisit
	}

	count := 0

	for y, row := range double {
		for x, ch := range row {
			if !visited[y][x] && (ch != 'X' && ch != 'S') {
				visited[y][y] = true
				count += floodFillDouble(double, visited, y, x)
			} else {
				visited[y][x] = true
			}
		}
	}

	return count
}

func floodFillDouble(double [][]rune, visited [][]bool, y, x int) int {
	isOnEdge := false
	count := 0
	q := [][]int{{y, x}}
	dirs := moves['S']
	h := len(double)
	w := len(double[0])

	for len(q) != 0 {
		current := q[0]
		q = dequeue(q)

		curY := current[0]
		curX := current[1]
		if curY == 0 || curY == h-1 || curX == 0 || curX == w-1 {
			isOnEdge = true
		}
		if !isOnEdge && double[curY][curX] != '0' && double[curY][curX] != 'S' {
			count++
		}

		for _, move := range dirs {
			nextY := move[0] + curY
			nextX := move[1] + curX

			if nextY >= 0 && nextY < h && nextX >= 0 && nextX < w && !visited[nextY][nextX] {
				nextCh := double[nextY][nextX]
				if nextCh != 'X' {
					visited[nextY][nextX] = true
					q = enqueue(q, []int{nextY, nextX})
				}
			}
		}
	}

	if isOnEdge {
		return 0
	} else {
		return count
	}
}
