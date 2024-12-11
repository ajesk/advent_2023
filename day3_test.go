package main

import (
	"testing"
)

func testScan(t *testing.T, grid []string, expected bool) {
	result := scanForSymbol(grid, 1, 1)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, !expected)
	}
}

func TestBasicScanForSymbol(t *testing.T) {
	test := []string{
		"...",
		".1.",
		"...",
	}
	testScan(t, test, false)
}

func TestScanForSymbolShouldNotBreakOnEdges(t *testing.T) {
	test := []string{
		"1",
	}

	// a failure is a panic here
	scanForSymbol(test, 0, 0)
}

func TestScanForSymbolShouldSuceedIfSymbolIsDirectlyToTheRight(t *testing.T) {
	test := []string{
		"...",
		".1&",
		"...",
	}

	testScan(t, test, true)
}

func TestScanForSymbolShouldSuceedIfSymbolIsOnTheBottom(t *testing.T) {
	test := []string{
		"...",
		".1&",
	}

	testScan(t, test, true)
}

func TestIsSymbolShouldReturnFalseIfNumber(t *testing.T) {
	test := '2'

	if isSymbol(test) {
		t.Errorf("Expected %v, got %v", false, true)
	}
}

func TestIsSymbolShouldReturnFalseIfPeriod(t *testing.T) {
	test := '.'

	if isSymbol(test) {
		t.Errorf("Expected %v, got %v", false, true)
	}
}

func TestIsSymbolShouldReturnFalseIfLetter(t *testing.T) {
	test := 'J'

	if isSymbol(test) {
		t.Errorf("Expected %v, got %v", false, true)
	}
}

func TestIsSymbolShouldReturnTrueIfValidSymbol(t *testing.T) {
	test := '%'

	if !isSymbol(test) {
		t.Errorf("Expected %v, got %v", true, false)
	}
}

func TestCountNumbersShouldReturnNothingWhenThereAreNoSymbols(t *testing.T) {
	test := []string{
		"..........",
		"...23.....",
		"..........",
	}
	expected := 0

	testFullCount(t, test, expected)
}

func TestCountNumbersShouldReturnOnlyValueWhenThereAreSymbolsTouching(t *testing.T) {
	test := []string{
		"..........",
		"...23.....",
		".....*....",
	}
	expected := 23

	testFullCount(t, test, expected)
}

func TestCountNumbersShouldReturnNoValueWhenThereAreSymbolsNotTouching(t *testing.T) {
	test := []string{
		"..........",
		"...23.....",
		"......*...",
	}
	expected := 0

	testFullCount(t, test, expected)
}

func TestCountNumbersShouldApplyIfLastNumber(t *testing.T) {
	test := []string{
		"..........",
		".........%",
		"........11",
	}
	expected := 11

	testFullCount(t, test, expected)
}

func TestAdventBase(t *testing.T) {
	test := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expected := 4361
	testFullCount(t, test, expected)
}

func testFullCount(t *testing.T, grid []string, expected int) {
	result := countNumbersWithAdjacentSymbols(grid)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay3Part1Output(t *testing.T) {
	t.Skip("day3 part 1 passed")
	result := dayThreePartOne()
	t.Errorf("day3, part1: %d", result)
}

// part 2
func TestFindAllSurroundingNumbersNoResults(t *testing.T) {
	test := []string{
		"...",
		".*.",
		"...",
	}
	expected := 0
	result := findAllSurroundingCells(test, 1, 1)

	if len(result) != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindAllSurroundingNumbersSingleResult(t *testing.T) {
	test := []string{
		"...",
		".*.",
		"4..",
	}
	expected := 1
	result := findAllSurroundingCells(test, 1, 1)

	if len(result) != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
	if result[0][0] != 2 || result[0][1] != 0 {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindAllSurroundingNumbersAllNumbers(t *testing.T) {
	test := []string{
		"111",
		"1*1",
		"411",
	}
	expected := 8
	result := findAllSurroundingCells(test, 1, 1)

	if len(result) != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindAllSurroundingNumbersOnCorner(t *testing.T) {
	test := []string{
		"...",
		"1..",
		"*1.",
	}
	expected := 2
	result := findAllSurroundingCells(test, 2, 0)

	if len(result) != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindAllSurroundingNumbersOnEdge(t *testing.T) {
	test := []string{
		".4.",
		"*..",
		"41.",
	}
	expected := 3
	result := findAllSurroundingCells(test, 1, 0)

	if len(result) != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindNumbersFromSurroundingCellsShouldBuildNothingButImmediateCell(t *testing.T) {
	test := []string{
		"....4....",
	}
	cells := [][]int{
		{0, 4},
	}
	expectedLen := 1
	expected := 4
	result := findNumbersFromSurroundingCells(test, cells)

	if len(result) != expectedLen {
		t.Errorf("Expected %v results, got %v", expectedLen, len(result))
	}

	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestFindNumbersToTheLeft(t *testing.T) {
	test := []string{
		"44444....",
	}
	cells := [][]int{
		{0, 4},
	}
	expectedLen := 1
	expected := 44444
	result := findNumbersFromSurroundingCells(test, cells)

	if len(result) != expectedLen {
		t.Errorf("Expected %v results, got %v", expectedLen, len(result))
	}

	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestFindNumbersToTheRight(t *testing.T) {
	test := []string{
		"....4444",
	}
	cells := [][]int{
		{0, 4},
	}
	expectedLen := 1
	expected := 4444
	result := findNumbersFromSurroundingCells(test, cells)

	if len(result) != expectedLen {
		t.Errorf("Expected %v results, got %v", expectedLen, len(result))
	}

	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestFindNumbersOnBothSides(t *testing.T) {
	test := []string{
		"...343..",
	}
	cells := [][]int{
		{0, 4},
	}
	expectedLen := 1
	expected := 343
	result := findNumbersFromSurroundingCells(test, cells)

	if len(result) != expectedLen {
		t.Errorf("Expected %v results, got %v", expectedLen, len(result))
	}

	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestFindNumbersShouldStopOnDot(t *testing.T) {
	test := []string{
		"..3.5.3.",
	}
	cells := [][]int{
		{0, 4},
	}
	expectedLen := 1
	expected := 5
	result := findNumbersFromSurroundingCells(test, cells)

	if len(result) != expectedLen {
		t.Errorf("Expected %v results, got %v", expectedLen, len(result))
	}

	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestFindNumbersShouldSplitResults(t *testing.T) {
	test := []string{
		"..3.53.",
	}
	cells := [][]int{
		{0, 2}, {0, 4},
	}
	expectedLen := 2
	result := findNumbersFromSurroundingCells(test, cells)

	if len(result) != expectedLen {
		t.Errorf("Expected %v results, got %v", expectedLen, len(result))
	}

	if result[0] != 3 || result[1] != 53 {
		t.Errorf("Expected values to match and they do not")
	}
}

func TestFindNumbersShouldNotClashIfMultipleCellsAreSideBySide(t *testing.T) {
	test := []string{
		"..359.",
	}
	cells := [][]int{
		{0, 2}, {0, 3},
	}
	expectedLen := 1
	expected := 359
	result := findNumbersFromSurroundingCells(test, cells)

	if len(result) != expectedLen {
		t.Errorf("Expected %v results, got %v", expectedLen, len(result))
	}

	if expected != result[0] {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestFindNumbersShouldGetResultsOnTwoLines(t *testing.T) {
	test := []string{
		"..359.",
		"12....",
	}
	cells := [][]int{
		{0, 2}, {1, 1},
	}
	expectedLen := 2
	result := findNumbersFromSurroundingCells(test, cells)

	if len(result) != expectedLen {
		t.Errorf("Expected %v results, got %v", expectedLen, len(result))
	}

	if result[0] != 359 || result[1] != 12 {
		t.Errorf("Expected values to match and they do not")
	}
}

func TestDay3Part2(t *testing.T) {
	t.Skip("day3 part 2 passed")
	result := dayThreePartTwo()
	t.Errorf("day3, part2: %d", result)

}
