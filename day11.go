package main

import (
	"math"
	"strings"
)

func day11Part1() int {
	fileContents, err := ReadFileLines("./data/day11.txt")
	print(fileContents)

	if err != nil {
		return 0
	}

	return calculateDistanceBetweenAllGalaxies(fileContents)
}

func day11Part2() int {
	fileContents, err := ReadFileLines("./data/day11.txt")
	print(fileContents)

	if err != nil {
		return 0
	}

	return alternateCalculateDistanceBetweenAllGalaxies(fileContents, 1000000)
}

func calculateDistanceBetweenAllGalaxies(universe []string) int {
	count := 0
	rowsToExpand := scanRowsForExpansion(universe)
	colsToExpand := scanColsForExpansion(universe)
	expandedUniverse := expandRows(universe, rowsToExpand)
	expandedUniverse = expandCols(expandedUniverse, colsToExpand)
	galaxyLocations := catalogGalaxies(expandedUniverse)

	for i, galA := range galaxyLocations {
		for j := i + 1; j < len(galaxyLocations); j++ {
			galB := galaxyLocations[j]
			count += calculateDistanceBetweenGalaxies(galA, galB)
		}
	}

	return count
}

func scanRowsForExpansion(universe []string) []bool {
	rowsToExpand := []bool{}
	for _, uniSlice := range universe {
		rowsToExpand = append(rowsToExpand, !strings.Contains(uniSlice, "#"))
	}
	return rowsToExpand
}

func scanColsForExpansion(universe []string) []bool {
	colsToExpand := []bool{}
	for x := range universe[0] {
		isEmpty := true
		for y := 0; y < len(universe); y++ {
			if universe[y][x] == '#' {
				isEmpty = false
			}
		}
		colsToExpand = append(colsToExpand, isEmpty)
	}
	return colsToExpand
}

func expandRows(universe []string, rowsToExpand []bool) []string {
	expandedUniverse := []string{}

	for i := len(universe) - 1; i >= 0; i-- {
		universeLine := universe[i]
		expandedUniverse = append([]string{universeLine}, expandedUniverse...)
		if rowsToExpand[i] {
			expandedUniverse = append([]string{universeLine}, expandedUniverse...)
		}
	}

	return expandedUniverse
}

func expandCols(universe []string, colsToExpand []bool) []string {
	expandedUniverse := []string{}

	for _, universeLine := range universe {
		expandedLine := universeLine
		for i := len(expandedLine) - 1; i >= 0; i-- {
			if colsToExpand[i] {
				expandedLine = expandedLine[:i] + "." + expandedLine[i:]
			}
		}
		expandedUniverse = append(expandedUniverse, expandedLine)
	}
	return expandedUniverse
}

func catalogGalaxies(universe []string) [][]int {
	galaxies := [][]int{}

	for y, row := range universe {
		for x, ch := range row {
			if ch == '#' {
				galaxies = append(galaxies, []int{y, x})
			}
		}
	}

	return galaxies
}

func calculateDistanceBetweenGalaxies(galA, galB []int) int {
	return int(math.Abs(float64(galB[0]-galA[0]))) + int(math.Abs(float64(galB[1]-galA[1])))
}

func howManyExpansionsBridged(galA, galB []int, expandedRows, expandedCols []bool) int {
	count := 0
	galAY := float64(galA[0])
	galBY := float64(galB[0])
	galAX := float64(galA[1])
	galBX := float64(galB[1])

	for i := math.Min(galAY, galBY); i < math.Max(galAY, galBY); i++ {
		if expandedRows[int(i)] {
			count++
		}
	}

	for i := math.Min(galAX, galBX); i < math.Max(galAX, galBX); i++ {
		if expandedCols[int(i)] {
			count++
		}
	}

	return count
}

func calculateDistanceBetweenGalaxiesFactor(galA, galB []int, expansionsBridged, factor int) int {
	yDif := int(math.Abs(float64(galB[0] - galA[0])))
	xDif := int(math.Abs(float64(galB[1] - galA[1])))
	factorCalculation := expansionsBridged*factor - expansionsBridged
	return yDif + xDif + factorCalculation
}

func alternateCalculateDistanceBetweenAllGalaxies(universe []string, expansionFactor int) int {
	count := 0
	rowsToExpand := scanRowsForExpansion(universe)
	colsToExpand := scanColsForExpansion(universe)
	galaxyLocations := catalogGalaxies(universe)

	for i, galA := range galaxyLocations {
		for j := i + 1; j < len(galaxyLocations); j++ {
			galB := galaxyLocations[j]
			expansionsBridged := howManyExpansionsBridged(galA, galB, rowsToExpand, colsToExpand)
			count += calculateDistanceBetweenGalaxiesFactor(galA, galB, expansionsBridged, expansionFactor)
		}
	}

	return count
}
