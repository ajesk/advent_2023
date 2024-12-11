package main

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type hand struct {
	cards []int
	wager int
	level int // one pair, 4 of a kind, etc
}

var cardConv = map[rune]int{
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}
var altCardConv = map[rune]int{
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func day7Part1() int {
	fileContents, err := ReadFileLines("./data/day7.txt")

	if err != nil {
		return 0
	}

	return getHandValues(fileContents, false)
}

func day7Part2() int {
	fileContents, err := ReadFileLines("./data/day7.txt")

	if err != nil {
		return 0
	}

	return getHandValues(fileContents, true)
}

func getHandValues(lines []string, jokers bool) int {
	hands := []hand{}
	total := 0
	for _, line := range lines {
		hands = append(hands, parseLineIntoHand(line, jokers))
	}

	sortHands(hands)

	for i := 0; i < len(hands); i++ {
		total += hands[i].wager * (len(hands) - i)
	}
	return total
}

func parseLineIntoHand(line string, jokers bool) hand {
	split := strings.Split(line, " ")
	wager, _ := strconv.Atoi(split[1])
	cards := parseCardValues(split[0], jokers)
	level := getHandLevel(cards, jokers)

	return hand{wager: wager, cards: cards, level: level}
}

func parseCardValues(cardString string, jokers bool) []int {
	cards := []int{}
	for _, card := range cardString {
		var val int

		if unicode.IsDigit(card) {
			val, _ = strconv.Atoi(string(card))
		} else {
			if jokers {
				val = altCardConv[card]
			} else {
				val = cardConv[card]
			}
		}

		cards = append(cards, val)
	}

	return cards
}

func getHandLevel(cards []int, jokers bool) int {
	cardCounts := map[int]int{}
	var actualCards []int

	if jokers {
		actualCards = buildJokerHand(cards)
	} else {
		actualCards = cards
	}

	for _, card := range actualCards {
		_, ok := cardCounts[card]
		if ok {
			cardCounts[card] = cardCounts[card] + 1
		} else {
			cardCounts[card] = 1
		}
	}

	unique := len(cardCounts)

	if unique == 1 { // 5 of a kind
		return 1
	} else if unique == 2 {
		for _, val := range cardCounts {
			if val == 4 || val == 1 {
				return 2
			} else {
				return 3
			}
		}
	} else if unique == 3 {
		for _, val := range cardCounts {
			if val == 3 {
				return 4
			}
		}
		return 5
	} else if unique == 4 {
		return 6
	}

	return 7
}

func buildJokerHand(orig []int) []int {
	cardCounts := map[int]int{}

	for _, card := range orig {
		_, ok := cardCounts[card]
		if ok {
			cardCounts[card] = cardCounts[card] + 1
		} else {
			cardCounts[card] = 1
		}
	}

	_, jokerExists := cardCounts[1]
	if !jokerExists {
		return orig
	}

	highCard := getHighestMostOccurringCard(cardCounts)

	if highCard == 1 {
		return orig
	}

	return replaceJokers(orig, highCard)
}

func replaceJokers(orig []int, highCard int) []int {
	newHand := make([]int, len(orig))
	copy(newHand, orig)

	for i := range newHand {
		if newHand[i] == 1 {
			newHand[i] = highCard
		}
	}

	return newHand
}

func getHighestMostOccurringCard(cardCounts map[int]int) int {
	var highCard int = 0
	var highCount int = 0

	for card, count := range cardCounts {
		if card == 1 {
			continue
		} else if count > highCount {
			highCard = card
			highCount = count
		} else if count == highCount && card > highCard {
			highCard = card
			highCount = count

		}
	}

	return highCard
}

func sortHands(hands []hand) {
	sort.Slice(hands, func(a, b int) bool {
		handA := hands[a]
		handB := hands[b]
		if handA.level != handB.level {
			return handA.level < handB.level
		}

		return secondarySort(handA.cards, handB.cards)
	})
}

func secondarySort(cardsA, cardsB []int) bool {
	for i := 0; i < len(cardsA); i++ {
		if cardsA[i] != cardsB[i] {
			return cardsA[i] > cardsB[i]
		}
	}
	return false
}
