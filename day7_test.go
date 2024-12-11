package main

import (
	"reflect"
	"testing"
)

func TestConvertLineToHand(t *testing.T) {
	test := "32T3K 765"
	expected := hand{
		cards: []int{3, 2, 10, 3, 13},
		wager: 765,
		level: 6,
	}
	result := parseLineIntoHand(test, false)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestParseCardStringShouldParseALineOfDigits(t *testing.T) {
	test := "32333"
	expected := []int{3, 2, 3, 3, 3}
	result := parseCardValues(test, false)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestParseCardStringShouldParseALineOfFaceCards(t *testing.T) {
	test := "TKQJA"
	expected := []int{10, 13, 12, 11, 14}
	result := parseCardValues(test, false)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHandLevelShouldResultInA5OfAKind(t *testing.T) {
	test := "TTTTT"
	expected := 1
	result := getHandLevel(parseCardValues(test, false), false)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHandLevelShouldResultInA4OfAKind(t *testing.T) {
	test := "TTTTK"
	expected := 2
	result := getHandLevel(parseCardValues(test, false), false)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHandLevelShouldResultInAFullHouse(t *testing.T) {
	test := "TKTTK"
	expected := 3
	result := getHandLevel(parseCardValues(test, false), false)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHandLevelShouldResultInAThreeOfAKind(t *testing.T) {
	test := "TKT5T"
	expected := 4
	result := getHandLevel(parseCardValues(test, false), false)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHandLevelShouldResultInATwoPair(t *testing.T) {
	test := "TK55T"
	expected := 5
	result := getHandLevel(parseCardValues(test, false), false)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHandLevelShouldResultInAPair(t *testing.T) {
	test := "7K55T"
	expected := 6
	result := getHandLevel(parseCardValues(test, false), false)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHandLevelShouldResultInAHand(t *testing.T) {
	test := "TK549"
	expected := 7
	result := getHandLevel(parseCardValues(test, false), false)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSortHandsByLevel(t *testing.T) {
	testA := hand{level: 4, cards: []int{1, 2, 3, 4, 5}}
	testB := hand{level: 1, cards: []int{2, 2, 3, 4, 5}}
	test := []hand{testA, testB}
	sortHands(test)

	if test[0].level != testB.level {
		t.Errorf("Expected %v, got %v", testB, test[0])
	}
}

func TestSortHandsByLevelInverse(t *testing.T) {
	testB := hand{level: 4, cards: []int{1, 2, 3, 4, 5}}
	testA := hand{level: 1, cards: []int{2, 2, 3, 4, 5}}
	test := []hand{testA, testB}
	sortHands(test)

	if test[0].level != testA.level {
		t.Errorf("Expected %v, got %v", testA, test[0])
	}
}

func TestSortHandsByHand(t *testing.T) {
	testB := hand{level: 1, cards: []int{1, 2, 3, 4, 5}}
	testA := hand{level: 1, cards: []int{2, 2, 3, 4, 5}}
	test := []hand{testA, testB}
	sortHands(test)

	if test[0].cards[0] != testA.cards[0] {
		t.Errorf("Expected %v, got %v", testA, test[0])
	}
}

func TestSortHandsByHandLastValue(t *testing.T) {
	testB := hand{level: 1, cards: []int{1, 2, 3, 4, 5}}
	testA := hand{level: 1, cards: []int{1, 2, 3, 4, 10}}
	test := []hand{testA, testB}
	sortHands(test)

	if test[0].cards[4] != testA.cards[4] {
		t.Errorf("Expected %v, got %v", testA, test[0])
	}
}

func TestSortHandsByHandAndLevel(t *testing.T) {
	testA := hand{level: 1, cards: []int{1, 2, 3, 4, 5}}
	testB := hand{level: 2, cards: []int{1, 2, 3, 4, 10}}
	testC := hand{level: 1, cards: []int{1, 2, 3, 4, 10}}
	test := []hand{testA, testB, testC}
	sortHands(test)

	if test[0].cards[4] != testC.cards[4] {
		t.Errorf("Expected %v, got %v", testC, test[0])
	}

	if test[2].level != testB.level {
		t.Errorf("Expected %v, got %v", testB, test[2])
	}
}

func TestDemoData(t *testing.T) {
	lines := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	expected := 6440
	result := getHandValues(lines, false)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReplaceJokers(t *testing.T) {
	test := []int{1, 1, 1, 1, 2}
	expected := [6]int{2, 2, 2, 2, 2}
	result := replaceJokers(test, 2)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReplaceNoJokers(t *testing.T) {
	test := []int{3, 2, 2, 2, 2}
	expected := [6]int{3, 2, 2, 2, 2}
	result := replaceJokers(test, 2)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReplaceAllJokers(t *testing.T) {
	test := []int{1, 1, 1, 1, 1}
	expected := [6]int{1, 1, 1, 1, 1}
	result := replaceJokers(test, 2)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReplaceSpecificJokers(t *testing.T) {
	test := []int{2, 2, 3, 1, 3}
	expected := [6]int{2, 2, 3, 3, 3}
	result := replaceJokers(test, 3)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHighestMostOccurringCardWhenAllTheSameOccurance(t *testing.T) {
	test := map[int]int{
		1: 1,
		2: 1,
		3: 1,
		4: 1,
		5: 1,
	}

	expected := 5
	result := getHighestMostOccurringCard(test)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHighestMostOccurringCardWhenTwoArdTheSame(t *testing.T) {
	test := map[int]int{
		1: 2,
		2: 2,
		3: 1,
	}

	expected := 2
	result := getHighestMostOccurringCard(test)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetHighestMostOccurringCardWhenOneIsHighest(t *testing.T) {
	test := map[int]int{
		1: 3,
		2: 1,
		3: 1,
	}

	expected := 3
	result := getHighestMostOccurringCard(test)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBuildJokerHandShouldReturnSameHandWithNoJoker(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	expected := [6]int{2, 3, 4, 5, 6}
	result := buildJokerHand(test)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBuildJokerHandShouldAPairWithOneJoker(t *testing.T) {
	test := []int{1, 3, 4, 5, 6}
	expected := [6]int{6, 3, 4, 5, 6}
	result := buildJokerHand(test)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBuildJokerHandShouldReturnThreeOfAKindWith2Jokers(t *testing.T) {
	test := []int{1, 3, 4, 5, 1}
	expected := [6]int{5, 3, 4, 5, 5}
	result := buildJokerHand(test)

	if reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDemoJokers(t *testing.T) {
	test, _ := ReadFileLines("./data/day7demo.txt")
	expected := 6839
	result := getHandValues(test, true)

	if expected != result {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
