package main

import (
	"testing"
)

func TestIsValidCounts(t *testing.T) {
	result := isValidColorCounts("red", 1)

	if !result {
		t.Errorf("Expected %v, got %v", !result, result)
	}

	result = isValidColorCounts("green", 1)

	if !result {
		t.Errorf("Expected %v, got %v", !result, result)
	}

	result = isValidColorCounts("blue", 1)

	if !result {
		t.Errorf("Expected %v, got %v", !result, result)
	}
}

func TestColorLimits(t *testing.T) {
	result := isValidColorCounts("red", 12)

	if !result {
		t.Errorf("Expected %v, got %v", !result, result)
	}

	result = isValidColorCounts("green", 13)

	if !result {
		t.Errorf("Expected %v, got %v", !result, result)
	}

	result = isValidColorCounts("blue", 14)

	if !result {
		t.Errorf("Expected %v, got %v", !result, result)
	}
}

func TestValidColorFail(t *testing.T) {
	result := isValidColorCounts("red", 13)

	if result {
		t.Errorf("Expected %v, got %v", !result, result)
	}

	result = isValidColorCounts("green", 14)

	if result {
		t.Errorf("Expected %v, got %v", !result, result)
	}

	result = isValidColorCounts("blue", 15)

	if result {
		t.Errorf("Expected %v, got %v", !result, result)
	}

	result = isValidColorCounts("notacolor", 1)

	if result {
		t.Errorf("Expected %v, got %v", !result, result)
	}
}

func TestRemovePrefixBasic(t *testing.T) {
	test := "Game 1: other shit"
	expected := "other shit"

	result := removePrefix(test)

	if result != expected {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}

func TestRemovePrefixShouldDoNothing(t *testing.T) {
	test := "same shit"
	expected := "same shit"

	result := removePrefix(test)

	if result != expected {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}

func TestValidateHandShouldSplitAndReturnSuccess(t *testing.T) {
	test := "1 green"
	expected := true
	result := validateHand(test)

	if result != expected {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}

func TestValidateHandShouldSplitAndReturnFailure(t *testing.T) {
	test := "14 green"
	expected := false
	result := validateHand(test)

	if result != expected {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}

func TestPowerOfGame(t *testing.T) {
	tests := map[string]int{
		"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green":                   48,
		"1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue":         12,
		"8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": 1560,
		"1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": 630,
		"6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green":                   36,
	}

	for test, expected := range tests {
		result := powerOfGame(test)

		if result != expected {
			t.Errorf("Expected '%v', got '%v'", expected, result)
		}
	}
}

func TestDayTwoPartOne(t *testing.T) {
	t.Skip("day2 part 1 passed")
	result := dayTwoPartOne()
	t.Logf("day2, part1: %d", result)
}

func TestDayTwoPartTwo(t *testing.T) {
	t.Skip("day2 part 2 passed")
	result := dayTwoPartTwo()
	t.Logf("day2, part2: %d", result)
}
