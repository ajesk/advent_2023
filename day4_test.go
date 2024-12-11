package main

import (
	"testing"
)

func TestStrip(t *testing.T) {
	test := "Card   1: 30 48 49 69 | 86 57 89"
	expected := "30 48 49 69 | 86 57 89"
	result := stripGame(test)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSimpleGame(t *testing.T) {
	winning := []string{"12", " 2", " 3", "42"}
	yours := []string{"12", " 2", " 5", "55", "60"}
	expected := 2

	result := countMatches(winning, yours)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay4Part1(t *testing.T) {
	// t.Skip("day4 part 1 passed")
	result := dayFourPartOne()
	t.Errorf("day4, part1: %d", result)
}

func TestHowManyGamesPlayed(t *testing.T) {
	test := []string{
		"Card   1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card   2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card   3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card   4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card   5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card   6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	expected := int64(30)

	result := howManyGamesPlayed(test)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay4Part2(t *testing.T) {
	// t.Skip("day4 part 2 passed answer: 11024379")
	result := dayFourPartTwo()
	t.Errorf("day4, part2: %d", result)
}
