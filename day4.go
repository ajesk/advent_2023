package main

import (
	"math"
	"strings"
)

func dayFourPartOne() int {
	fileContents, err := ReadFileLines("./data/day4.txt")

	if err != nil {
		return 0
	}

	return checkGamesForWinners(fileContents)
}

func checkGamesForWinners(games []string) int {
	count := 0

	for _, game := range games {
		strippedGame := stripGame(game)
		split := strings.Split(strippedGame, " | ")
		winningNumbers := strings.Split(split[0], " ")
		yourNumbers := strings.Split(split[1], " ")
		matches := float64(countMatches(winningNumbers, yourNumbers))
		if matches != 0 {
			count += int(math.Pow(2, matches-1))
		}
	}

	return count
}

func dayFourPartTwo() int64 {
	fileContents, err := ReadFileLines("./data/day4.txt")

	if err != nil {
		return 0
	}

	return howManyGamesPlayed(fileContents)
}

func howManyGamesPlayed(games []string) int64 {
	tab := make([]int64, len(games))
	var matchesList = getMatchesList(games)
	count := int64(0)

	for i := range tab {
		tab[i] = 1
	}

	for i := range tab {
		count += tab[i]
		if i == len(tab)-1 {
			break
		}
		currentWins := matchesList[i]
		for x := 1; x < currentWins+1 && x+i < len(tab); x++ {
			tab[x+i] = tab[x+i] + tab[i]
		}
	}

	return count
}

func getMatchesList(games []string) []int {
	var matches = []int{}
	for _, game := range games {
		currentGame := stripGame(game)
		splitGame := strings.Split(currentGame, " | ")
		winningNumbers := strings.Split(splitGame[0], " ")
		yourNumbers := strings.Split(splitGame[1], " ")
		matchCount := countMatches(winningNumbers, yourNumbers)
		matches = append(matches, matchCount)
	}

	return matches

}

func stripGame(txt string) string {
	return txt[10:]
}

func countMatches(winningNumbers []string, yourNumbers []string) int {
	count := 0

	for _, winNum := range winningNumbers {
		if winNum == "" {
			continue
		}
		for _, yourNum := range yourNumbers {
			if yourNum == "" {
				continue
			}
			if strings.Trim(winNum, " ") == strings.Trim(yourNum, " ") {
				count++
			}
		}
	}
	return count
}
