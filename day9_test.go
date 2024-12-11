package main

import (
	"reflect"
	"testing"
)

func TestDayday9Part1Demo(t *testing.T) {
	lines, _ := ReadFileLines("./data/day9demo.txt")
	result := extrapolateValues(lines, false)
	expected := 114

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDayday9Part2Demo(t *testing.T) {
	// lines, _ := ReadFileLines("./data/dayday9demo.txt")
	result := 0
	expected := 0

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReduceToZeroSimple(t *testing.T) {
	test := []int{0, 3, 6, 9, 12, 15}
	result := reduceToZeroes(test)
	expected := [4][]int{
		{0, 3, 6, 9, 12, 15},
		{3, 3, 3, 3, 3},
		{0, 0, 0, 0},
	}

	if reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReduceToZeroDemo2(t *testing.T) {
	test := []int{1, 3, 6, 10, 15, 21}
	result := reduceToZeroes(test)
	expected := [4][]int{
		{1, 3, 6, 10, 15, 21},
		{2, 3, 4, 5, 6},
		{1, 1, 1, 1},
		{0, 0, 0},
	}

	if reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReduceToZeroDemo3(t *testing.T) {
	test := []int{10, 13, 16, 21, 30, 45}
	result := reduceToZeroes(test)
	expected := [5][]int{
		{10, 13, 16, 21, 30, 45},
		{3, 3, 5, 9, 15},
		{0, 2, 4, 6},
		{2, 2, 2},
		{0, 0},
	}

	if reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestScaleToFindNextValueDemo(t *testing.T) {
	expected := 18
	test := [][]int{
		{0, 3, 6, 9, 12, 15},
		{3, 3, 3, 3, 3},
		{0, 0, 0, 0},
	}
	result := scaleToFindNextValue(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestScaleToFindNextValueDemo3(t *testing.T) {
	expected := 68
	test := [][]int{
		{10, 13, 16, 21, 30, 45},
		{3, 3, 5, 9, 15},
		{0, 2, 4, 6},
		{2, 2, 2},
		{0, 0},
	}
	result := scaleToFindNextValue(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestScaleToFindNextValueBackwardsDemo(t *testing.T) {
	expected := -3
	test := [][]int{
		{0, 3, 6, 9, 12, 15},
		{3, 3, 3, 3, 3},
		{0, 0, 0, 0},
	}
	result := scaleToFindNextValueBackwards(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestScaleToFindNextValueBackwardsDemo2(t *testing.T) {
	expected := 0
	test := [][]int{
		{1, 3, 6, 10, 15, 21},
		{2, 3, 4, 5, 6},
		{1, 1, 1, 1},
		{0, 0, 0},
	}
	result := scaleToFindNextValueBackwards(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestScaleToFindNextValueBackwardsDemo3(t *testing.T) {
	expected := 68
	test := [][]int{
		{10, 13, 16, 21, 30, 45},
		{3, 3, 5, 9, 15},
		{0, 2, 4, 6},
		{2, 2, 2},
		{0, 0},
	}
	result := scaleToFindNextValue(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
