package main

import (
	"reflect"
	"testing"
)

func TestConvertLineToMap(t *testing.T) {
	test := "12 13 14"
	expected := mapping{source: 13, dest: 12, rng: 14, sourceEnd: 27, destEnd: 26, diff: -1}

	result := lineToMap(test)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetSeeds(t *testing.T) {
	test := "seeds: 20 23 24"
	expected := [4]int{20, 23, 24}

	result := getSeeds(test)
	if reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetSeedRanges(t *testing.T) {
	test := "seeds: 79 14 55 13"
	expected := []seedRange{
		{start: 79, end: 93},
		{start: 55, end: 68},
	}

	result := getSeedRanges(test)
	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if result[0].end != expected[0].end {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}

func TestBuildAlmanacShouldEnterASingleVal(t *testing.T) {
	test := []string{
		"12 13 14",
	}
	result := convertLinesIntoAlmanac(test)
	expectedGroups := 1
	expectedEntries := 1

	if len(result) != expectedGroups {
		t.Errorf("Expected %v, got %v", expectedGroups, len(result))
	}

	if len(result[0]) != expectedEntries {
		t.Errorf("Expected %v, got %v", expectedGroups, len(result[0]))
	}
}

func TestBuildAlmanacShouldEnterTwoVals(t *testing.T) {
	test := []string{
		"12 13 14",
		"5 2 1",
	}
	result := convertLinesIntoAlmanac(test)
	expectedGroups := 1
	expectedEntries := 2

	if len(result) != expectedGroups {
		t.Errorf("Expected %v, got %v", expectedGroups, len(result))
	}

	if len(result[0]) != expectedEntries {
		t.Errorf("Expected %v, got %v", expectedEntries, len(result[0]))
	}
}

func TestBuildAlmanacShouldEnterDiffValsInTwoGroups(t *testing.T) {
	test := []string{
		"12 13 14",
		"",
		"4 4 4",
		"4 4 4",
		"4 4 4",
	}
	result := convertLinesIntoAlmanac(test)
	expectedGroups := 2
	expectedEntries1 := 1
	expectedEntries2 := 3

	if len(result) != expectedGroups {
		t.Errorf("Expected %v, got %v", expectedGroups, len(result))
	}

	if len(result[0]) != expectedEntries1 {
		t.Errorf("Expected %v, got %v", expectedEntries1, len(result[0]))
	}

	if len(result[1]) != expectedEntries2 {
		t.Errorf("Expected %v, got %v", expectedEntries2, len(result[1]))
	}
}

func TestBuildAlmanacShouldIgnoreFluffLines(t *testing.T) {
	test := []string{
		"12 13 14",
		"",
		"Seed to soil mapping",
		"4 4 4",
		"4 4 4",
		"4 4 4",
	}
	result := convertLinesIntoAlmanac(test)
	expectedGroups := 2
	expectedEntries1 := 1
	expectedEntries2 := 3

	if len(result) != expectedGroups {
		t.Errorf("Expected %v, got %v", expectedGroups, len(result))
	}

	if len(result[0]) != expectedEntries1 {
		t.Errorf("Expected %v, got %v", expectedEntries1, len(result[0]))
	}

	if len(result[1]) != expectedEntries2 {
		t.Errorf("Expected %v, got %v", expectedEntries2, len(result[1]))
	}
}

func TestTransferLayerShouldApplyPositiveDiff(t *testing.T) {
	test := []string{
		"22 12 4",
	}
	mappings := convertLinesIntoAlmanac(test)
	expected := 23

	result := transferLayer(13, mappings[0])

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTransferLayerShouldApplyNegativeDiff(t *testing.T) {
	test := []string{
		"12 22 4",
	}
	mappings := convertLinesIntoAlmanac(test)
	expected := 13

	result := transferLayer(23, mappings[0])

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTransferLayerShouldReturnSameValIfNotInMapping(t *testing.T) {
	test := []string{
		"12 22 4",
	}
	mappings := convertLinesIntoAlmanac(test)
	expected := 33

	result := transferLayer(expected, mappings[0])

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDemoSeedValidation(t *testing.T) {
	result := day5Part1File("./data/day5demo.txt")
	expected := 35

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSeedValidation(t *testing.T) {
	result := day5Part1File("./data/day5.txt")
	expected := 57075758

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSeedShouldBeInRangeSimple(t *testing.T) {
	seedRng := []seedRange{
		{start: 1, end: 3},
	}
	result := isSeedInRange(2, seedRng)

	if !result {
		t.Errorf("Expected %v, got %v", !result, result)
	}
}

func TestSeedShouldBeInRangeOnStart(t *testing.T) {
	seedRng := []seedRange{
		{start: 1, end: 3},
	}
	result := isSeedInRange(1, seedRng)

	if !result {
		t.Errorf("Expected %v, got %v", !result, result)
	}
}

func TestSeedShouldBeInRangeOnEnd(t *testing.T) {
	seedRng := []seedRange{
		{start: 1, end: 3},
	}
	result := isSeedInRange(3, seedRng)

	if !result {
		t.Errorf("Expected %v, got %v", !result, result)
	}
}

func TestSeedShouldBeOutOfRange(t *testing.T) {
	seedRng := []seedRange{
		{start: 1, end: 3},
	}
	result := isSeedInRange(4, seedRng)

	if result {
		t.Errorf("Expected %v, got %v", !result, result)
	}
}

func TestSeedShouldBeOutOfRangeBefore(t *testing.T) {
	seedRng := []seedRange{
		{start: 2, end: 4},
	}
	result := isSeedInRange(1, seedRng)

	if result {
		t.Errorf("Expected %v, got %v", !result, result)
	}
}

func TestSeedShouldBeOutOfRangeBetween(t *testing.T) {
	seedRng := []seedRange{
		{start: 2, end: 4},
		{start: 6, end: 10},
	}
	result := isSeedInRange(5, seedRng)

	if result {
		t.Errorf("Expected %v, got %v", !result, result)
	}
}

func TestReverseLookupBase(t *testing.T) {
	test := []string{
		"12 13 14",
		"",
		"13 14 4",
	}
	almanac := convertLinesIntoAlmanac(test)
	result := reverseLookupSource(13, almanac)
	expected := 15

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDemoSeedRanges(t *testing.T) {
	result := day5Part2File("./data/day5demo.txt")
	expected := 46

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
