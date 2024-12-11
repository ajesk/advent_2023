package main

import (
	"testing"
)

func TestDay14Part1Demo(t *testing.T) {
	lines, _ := ReadFileLines("./data/day14demo.txt")
	result := rollAndCount(lines)
	expected := 136

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay14Part2Demo(t *testing.T) {
	lines, _ := ReadFileLines("./data/day14demo.txt")
	result := rollAndCount(lines)
	expected := 0

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
