package main

import (
	"log"
	"time"
)

type problem struct {
	name   string
	day    int
	num    int
	method func() int
}

func main() {
	problems := []problem{}

	// problems = append(problems, buildProblem("Trebuchet?!", 1, 1, dayOnePartOne))
	// problems = append(problems, buildProblem("Calibration", 1, 2, dayOnePartTwo))
	// problems = append(problems, buildProblem("Cube Conundrum", 2, 1, dayTwoPartOne))
	// problems = append(problems, buildProblem("Fewest Cubes", 2, 2, dayTwoPartTwo))
	// problems = append(problems, buildProblem("Gear Ratios", 3, 1, dayThreePartOne))
	// problems = append(problems, buildProblem("The Actual Gears", 3, 2, dayThreePartTwo))
	// problems = append(problems, buildProblem("Scratchcards", 4, 1, dayFourPartOne))
	// problems = append(problems, buildProblem("Exponential Scratchcards", 4, 2, dayFourPartTwo))
	// problems = append(problems, buildProblem("Min of Seeds", 5, 1, day5Part1))
	// problems = append(problems, buildProblem("Min of Seed Ranges", 5, 2, day5Part2))
	// problems = append(problems, buildProblem("Toy Boat Race", 6, 1, day6Part1))
	// problems = append(problems, buildProblem("BIG Toy Boat Race", 6, 2, day6Part2))
	// problems = append(problems, buildProblem("Poker Face", 7, 1, day7Part1))
	// problems = append(problems, buildProblem("Joker Face", 7, 2, day7Part2))
	// problems = append(problems, buildProblem("Navigating the Wasteland", 8, 1, day8Part1))
	// problems = append(problems, buildProblem("Navigating the Ghost Wasteland", 8, 2, day8Part2))
	// problems = append(problems, buildProblem("Mirage Maintenance", 9, 1, day9Part1))
	// problems = append(problems, buildProblem("Mirage Maintenance Backwards", 9, 2, day9Part2))
	// problems = append(problems, buildProblem("Pipe Maze", 10, 1, day10Part1))
	// problems = append(problems, buildProblem("Pipe Maze Gaps", 10, 2, day10Part2))
	// problems = append(problems, buildProblem("Cosmic Expansion", 11, 1, day11Part1))
	// problems = append(problems, buildProblem("MASSIVE Cosmic Expansion", 11, 2, day11Part2))
	problems = append(problems, buildProblem("Roll Some Rocks", 14, 1, day14Part1))
	runResults(problems)
}

func buildProblem(name string, day int, num int, method func() int) problem {
	return problem{
		name:   name,
		day:    day,
		num:    num,
		method: method,
	}
}

func runResults(problems []problem) {

	for _, prob := range problems {
		start := time.Now()
		// result := "redacted"
		result := prob.method()
		duration := time.Since(start)

		log.Printf(
			"[Day %v, Part %v, %v]: result = %v, duration = %v",
			prob.day, prob.num, prob.name, result, duration,
		)
	}
}
