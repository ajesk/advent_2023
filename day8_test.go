package main

import (
	"testing"
)

func TestDay8Part1(t *testing.T) {
	// t.Skip("day4 part 1 passed")
	result := dayFourPartOne()
	t.Errorf("day4, part1: %d", result)
}

func TestParsePaths(t *testing.T) {
	test := []string{
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
	}

	result := parsePaths(test)
	first := result["AAA"]
	if len(result) != 2 {
		t.Errorf("Expected %v, got %v", 2, result)
	}

	if first[0] != "BBB" || first[1] != "CCC" {
		t.Errorf("Expected %v, got %v", first, result)
	}
}

func TestDemo1(t *testing.T) {
	lines, _ := ReadFileLines("./data/day8demo.txt")
	result := navigateWasteland(lines)
	expected := 2

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDemo2(t *testing.T) {
	lines := []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	result := navigateWasteland(lines)
	expected := 6

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGhostDemo(t *testing.T) {
	lines, _ := ReadFileLines("./data/day8demo2.txt")
	result := navigateGhostWasteland(lines)
	expected := 6

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
