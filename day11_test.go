package main

import (
	"reflect"
	"testing"
)

func TestDay11Part1Demo(t *testing.T) {
	lines, _ := ReadFileLines("./data/day11demo.txt")
	result := calculateDistanceBetweenAllGalaxies(lines)
	expected := 374

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay11Part2Demo(t *testing.T) {
	lines, _ := ReadFileLines("./data/day11demo.txt")
	result := alternateCalculateDistanceBetweenAllGalaxies(lines, 10)
	expected := 1030

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay11Part2Demo2(t *testing.T) {
	lines, _ := ReadFileLines("./data/day11demo.txt")
	result := alternateCalculateDistanceBetweenAllGalaxies(lines, 100)
	expected := 8410

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestScanRowsForExpansionSingleRowExpand(t *testing.T) {
	test := [][]string{
		{"........."},
		{".....#..."},
		{".........", ".....#...."},
		{".........", ".........."},
		{"..#......", ".....#...."},
	}
	expected := [][]bool{
		{true},
		{false},
		{true, false},
		{true, true},
		{false, false}}

	for i, galaxy := range test {
		result := scanRowsForExpansion(galaxy)
		expect := expected[i]

		if !reflect.DeepEqual(result, expect) {
			t.Errorf("Expected %v, got %v", expect, result)
		}
	}

}

func TestScanColsForExpansionSingleRowExpand(t *testing.T) {
	test := [][]string{
		{".", "."},
		{".", "#"},
		{"..", ".#"},
		{"..", ".."},
		{"#.", ".#"},
	}
	expected := [][]bool{
		{true},
		{false},
		{true, false},
		{true, true},
		{false, false},
	}

	for i, galaxy := range test {
		result := scanColsForExpansion(galaxy)
		expect := expected[i]

		if !reflect.DeepEqual(result, expect) {
			t.Errorf("Expected %v, got %v", expect, result)
		}
	}
}

func TestExpandRows(t *testing.T) {
	test := [][]string{
		{"....."},
		{"..#.."},
		{".....", ".#..."},
		{".....", "....."},
		{"..#..", ".#..."},
		{".....", ".#...", "....."},
	}
	expected := [][]string{
		{".....", "....."},
		{"..#.."},
		{".....", ".....", ".#..."},
		{".....", ".....", ".....", "....."},
		{"..#..", ".#..."},
		{".....", ".....", ".#...", ".....", "....."},
	}

	for i, testGal := range test {
		rowsToExpand := scanRowsForExpansion(testGal)
		result := expandRows(testGal, rowsToExpand)
		expect := expected[i]

		if !reflect.DeepEqual(result, expect) {
			t.Errorf("Expected %v, got %v", expect, result)
		}
	}
}

func TestExpandCols(t *testing.T) {
	test := [][]string{
		{".", "."},
		{".", "#"},
		{"..", ".#"},
		{"..", ".."},
		{"#.", ".#"},
	}
	expected := [][]string{
		{"..", ".."},
		{".", "#"},
		{"...", "..#"},
		{"....", "...."},
		{"#.", ".#"},
	}

	for i, testGal := range test {
		colsToExpand := scanColsForExpansion(testGal)
		result := expandCols(testGal, colsToExpand)
		expect := expected[i]

		if !reflect.DeepEqual(result, expect) {
			t.Errorf("Expected %v, got %v", expect, result)
		}
	}
}

func TestCatalogGalaxies(t *testing.T) {
	test := []string{
		"...#.",
		".#...",
		"#....",
	}
	expected := [][]int{
		{0, 3}, {1, 1}, {2, 0},
	}

	result := catalogGalaxies(test)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCalculateDistanceSameLine(t *testing.T) {
	tests := [][][]int{
		{{1, 0}, {1, 5}},
		{{6, 1}, {1, 5}},
		{{2, 0}, {10, 0}},
		{{2, 0}, {7, 12}},
		{{1, 9}, {2, 0}},
	}
	expected := []int{
		5,
		9,
		8,
		17,
		10,
	}
	for i, test := range tests {

		result := calculateDistanceBetweenGalaxies(test[0], test[1])
		if result != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], result)
		}
	}
}

func TestHowManyExpansionsBridged(t *testing.T) {
	test, _ := ReadFileLines("./data/day11demo.txt")
	expandedRows := scanRowsForExpansion(test)
	expandedCols := scanColsForExpansion(test)
	cases := [][][]int{
		{{0, 3}, {1, 7}},
		{{1, 7}, {2, 0}},
		{{2, 0}, {4, 6}},
		{{0, 3}, {8, 7}},
	}
	expected := []int{
		1,
		2,
		3,
		3,
	}
	for i, testCase := range cases {

		result := howManyExpansionsBridged(testCase[0], testCase[1], expandedRows, expandedCols)
		if result != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], result)
		}
	}
}
