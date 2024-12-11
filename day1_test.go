package main

import (
	"testing"
)

func TestRemoveAllNumbersWithEdges(t *testing.T) {
	test := "1jdfadjlkfd2"
	expected := "12"

	result := removeAllNonNumbers(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRemoveAllNumbersWithEmbedded(t *testing.T) {
	test := "gjdf1djl2fdf"
	expected := "12"

	result := removeAllNonNumbers(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRemoveAllNumbersWithSameNumber(t *testing.T) {
	test := "gjdfdjl2fdf"
	expected := "2"

	result := removeAllNonNumbers(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRemoveAllNumbersWithExtraNumbers(t *testing.T) {
	test := "gjd2d3l2fdf"
	expected := "232"

	result := removeAllNonNumbers(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRemoveAllNumbersWithNoNumbers(t *testing.T) {
	test := "gjddjlfdf"
	expected := ""

	result := removeAllNonNumbers(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetNumberOfLine(t *testing.T) {
	test := "gjd2djl2fdf"
	expected := 22

	result := getNumberOfLine(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetSumOfAllLines(t *testing.T) {
	test := [2]string{"gjd2djl2fdf", "2kjfasd1"}
	expected := 43

	result := getSumOfAllLines(test[:])

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReplaceFirstBase(t *testing.T) {
	test := "two"
	expected := "2"

	result := replaceFirstNumber(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFirstNumberShouldReplaceOnlyFirst(t *testing.T) {
	test := "onetwothree"
	expected := "1twothree"

	result := replaceFirstNumber(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReplaceFirstNumberFailure(t *testing.T) {
	test := "bhfhszrhzgrhsfd2threeseventwosevenoneseven"
	expected := "bhfhszrhzgrhsfd23seventwosevenoneseven"

	result := replaceFirstNumber(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReplaceLastBase(t *testing.T) {
	test := "two"
	expected := "2"

	result := replaceLastNumber(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReplaceLastShouldOnlyReplaceLast(t *testing.T) {
	test := "onetwothree"
	expected := "onetwo3"

	result := replaceLastNumber(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReplaceLastNumberFailure(t *testing.T) {
	test := "bhfhszrhzgrhsfd2threeseventwosevenoneseven"
	expected := "bhfhszrhzgrhsfd2threeseventwosevenone7"

	result := replaceLastNumber(test)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPart1Output(t *testing.T) {
	t.Skip("day1 part 1 passed")
	result := dayOnePartOne()
	t.Logf("day1, part1: %d", result)
}

func TestPart2Output(t *testing.T) {
	t.Skip("day1 part 2 passed")
	result := dayOnePartTwo()
	t.Logf("day1, part2: %d", result)
}
