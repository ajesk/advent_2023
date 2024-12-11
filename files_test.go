package main

import (
	"reflect"
	"testing"
)

func TestReadFileLines(t *testing.T) {
	fileContents := []string{
		"line 1",
		"line 2",
	}

	fileName := "./data/test.txt"

	result, err := ReadFileLines(fileName)
	if err != nil {
		t.Fatalf("Error reading file: %s", err)
	}

	if !reflect.DeepEqual(result, fileContents) {
		t.Errorf("Expected %v, got %v", fileContents, result)
	}
}
