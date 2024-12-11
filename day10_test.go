package main

import (
	"reflect"
	"testing"
)

func TestDay10Part1Demo(t *testing.T) {
	lines, _ := ReadFileLines("./data/day10demo.txt")
	result, _ := bfs(lines)
	expected := 4

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay10Part2Demo(t *testing.T) {
	result := day10Part2File("./data/day10demo.txt")
	expected := 1

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay10Part2Demo2(t *testing.T) {
	result := day10Part2File("./data/day10demo2.txt")
	expected := 1

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay10Part2Demo3(t *testing.T) {
	result := day10Part2File("./data/day10demo3.txt")
	expected := 4

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay10Part2Demo4(t *testing.T) {
	result := day10Part2File("./data/day10demo4.txt")
	expected := 8

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay10Part2Demo5(t *testing.T) {
	result := day10Part2File("./data/day10demo5.txt")
	expected := 10

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay10Part1(t *testing.T) {
	day10Part1()
}

func TestDay10Part1Demo2(t *testing.T) {
	lines, _ := ReadFileLines("./data/day10demo2.txt")
	result, _ := bfs(lines)
	expected := 8

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindTheStart(t *testing.T) {
	lines, _ := ReadFileLines("./data/day10demo.txt")
	result := findStartS(lines)
	expected := []int{1, 1}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindTheStart2(t *testing.T) {
	lines, _ := ReadFileLines("./data/day10demo2.txt")
	result := findStartS(lines)
	expected := []int{2, 0}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCreateDoubleShouldHaveDoubledTheSize(t *testing.T) {
	lines, _ := ReadFileLines("./data/day10demo.txt")
	result := createDoubleArray(lines)
	expectedHeight := len(lines) * 2
	expectedWidth := len(lines[0]) * 2

	if expectedHeight != len(result) {
		t.Errorf("Expected %v, got %v", expectedHeight, result)
	}

	if expectedWidth != len(result[0]) {
		t.Errorf("Expected %v, got %v", expectedHeight, result)
	}
	expectedWidth = len(lines[1]) * 2
	if expectedWidth != len(result[1]) {
		t.Errorf("Expected %v, got %v", expectedHeight, result)
	}
}

func TestCreateDoubleShouldHaveEveryOtherLineBeEmpty(t *testing.T) {
	lines, _ := ReadFileLines("./data/day10demo.txt")
	result := createDoubleArray(lines)[1]

	for _, ch := range result {
		if ch != '0' {
			t.Errorf("Expected %v, got %v", "0", ch)
		}
	}
}

func TestCreateDoubleShouldHaveEveryRunBeEmptyInRow(t *testing.T) {
	lines, _ := ReadFileLines("./data/day10demo.txt")
	result := createDoubleArray(lines)[0]

	for i, ch := range result {
		if i%2 == 1 && ch != '0' {
			t.Errorf("Expected %v, got %v", "0", ch)
		}
	}
	result = createDoubleArray(lines)[2]

	for i, ch := range result {
		if i%2 == 1 && ch != '0' {
			t.Errorf("Expected %v, got %v", "0", ch)
		}
	}
}

func TestCreateDoubleShouldHaveCopiedValues(t *testing.T) {
	lines, _ := ReadFileLines("./data/day10demo.txt")
	line := lines[0]
	result := createDoubleArray(lines)
	resultLine := result[0]

	for i, ch := range line {
		if ch != resultLine[i*2] {
			t.Errorf("Expected %v, got %v", ch, resultLine[i*2])
		}
	}
	line = lines[2]
	resultLine = result[4]

	for i, ch := range line {
		if ch != resultLine[i*2] {
			t.Errorf("Expected %v, got %v", ch, resultLine[i*2])
		}
	}
}
