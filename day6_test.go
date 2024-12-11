package main

import (
	"testing"
)

func TestParseDataLine(t *testing.T) {
	test := []string{
		"Time:     31       18      13     56",
	}

	expected := []int{31, 18, 13, 56}

	result := parseLineNumbers(test[0])
	if len(result) != len(expected) && expected[0] == result[0] {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindLowerLimit(t *testing.T) {
	result := findLowerLimit(7, 9)
	expected := 2

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindUpperLimit(t *testing.T) {
	result := findUpperLimit(7, 9)
	expected := 5

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindRange(t *testing.T) {
	result := findRecordRange(7, 9)
	expected := 4

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDemoDataPart1(t *testing.T) {
	file, _ := ReadFileLines("./data/day6demo.txt")
	expected := 288
	result := compileRecordRanges(file)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestStripNumbers(t *testing.T) {
	test := []string{
		"Time:     31       18      13     56",
	}

	result := squishNumber(test[0])
	expected := 31181356
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCompileLargeRac(t *testing.T) {
	file, _ := ReadFileLines("./data/day6demo.txt")
	expected := 71503
	result := compileLargeRecord(file)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
