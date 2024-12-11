package main

import (
	"testing"
)

func TestDay12Part1Demo(t *testing.T) {
	// lines, _ := ReadFileLines("./data/day12demo.txt")
	result := 0
	expected := 0

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay12Part2Demo(t *testing.T) {
	// lines, _ := ReadFileLines("./data/day12demo.txt")
	result := 0
	expected := 0

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBreakDownSpringCounts(t *testing.T) {
	cases := []string{
		"???.###",
		".??..??...?##.",
		"????.#...#...",
		"????.######..#####.",
	}

	print(cases)
}
