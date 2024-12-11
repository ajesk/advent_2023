package main

import (
	"strconv"
	"strings"
	"unicode"
)

type mapping struct {
	source    int
	dest      int
	rng       int
	sourceEnd int
	destEnd   int
	diff      int
}

type seedRange struct {
	start int
	end   int
}

func day5Part1() int {
	return day5Part1File("./data/day5.txt")
}

func day5Part1File(filename string) int {
	fileContents, err := ReadFileLines(filename)

	if err != nil {
		return 0
	}

	seeds := getSeeds(fileContents[0])
	almanac := convertLinesIntoAlmanac(fileContents[2:])

	plantSeeds(&seeds, almanac)

	var minSeed int
	for i, seed := range seeds {
		if i == 0 || seed < minSeed {
			minSeed = seed
		}
	}
	return minSeed

}

func day5Part2() int {
	return day5Part2File("./data/day5.txt")
}

func day5Part2File(filename string) int {
	fileContents, err := ReadFileLines(filename)

	if err != nil {
		return 0
	}

	seeds := getSeedRanges(fileContents[0])
	almanac := convertLinesIntoAlmanac(fileContents[2:])

	return reverseLookupSeedRange(seeds, almanac)
}

func plantSeeds(seeds *[]int, almanac [][]mapping) {
	for _, almanacMap := range almanac {
		for i := 0; i < len(*seeds); i++ {
			currentSeed := (*seeds)[i]
			(*seeds)[i] = transferLayer(currentSeed, almanacMap)
		}
	}
}

func getSeeds(line string) []int {
	seeds := []int{}
	lineVals := strings.Split(line, " ")[1:]

	for _, seed := range lineVals {
		seedVal, _ := strconv.Atoi(seed)
		seeds = append(seeds, seedVal)
	}

	return seeds
}

func transferLayer(value int, maps []mapping) int {
	for _, mapEntry := range maps {
		if value >= mapEntry.source && value <= mapEntry.sourceEnd {
			return value + mapEntry.diff
		}
	}

	return value
}

func convertLinesIntoAlmanac(lines []string) [][]mapping {
	almanac := [][]mapping{}
	level := []mapping{}

	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			almanac = append(almanac, level)
			level = []mapping{}
		} else if unicode.IsDigit(rune(line[0])) {
			level = append(level, lineToMap(line))
		}
	}
	almanac = append(almanac, level)
	return almanac
}

func lineToMap(line string) mapping {
	vals := strings.Split(line, " ")
	source, _ := strconv.Atoi(vals[1])
	dest, _ := strconv.Atoi(vals[0])
	rng, _ := strconv.Atoi(vals[2])
	sourceEnd := source + rng
	destEnd := dest + rng
	diff := dest - source
	return mapping{
		source:    source,
		dest:      dest,
		rng:       rng,
		sourceEnd: sourceEnd,
		destEnd:   destEnd,
		diff:      diff,
	}
}

// part 2 addition
func getSeedRanges(line string) []seedRange {
	seeds := []seedRange{}
	lineVals := strings.Split(line, " ")[1:]

	for i := 0; i < len(lineVals); i += 2 {
		start, _ := strconv.Atoi(lineVals[i])
		number, _ := strconv.Atoi(lineVals[i+1])
		end := start + number

		seeds = append(seeds, seedRange{start: start, end: end})
	}

	return seeds
}

func isSeedInRange(seed int, sdRange []seedRange) bool {
	for _, rng := range sdRange {
		if seed >= rng.start && seed <= rng.end {
			return true
		}
	}
	return false
}

func reverseLookupSource(value int, almanac [][]mapping) int {
	runningValue := value

	for i := len(almanac) - 1; i >= 0; i-- {
		currentMapping := almanac[i]

		for _, mapEntry := range currentMapping {
			if runningValue >= mapEntry.dest && runningValue <= mapEntry.destEnd {
				runningValue = runningValue - mapEntry.diff
				break
			}
		}
	}

	return runningValue

}

func reverseLookupSeedRange(seedRanges []seedRange, almanac [][]mapping) int {
	for i := 0; i < 100000000; i++ {
		sourceSeed := reverseLookupSource(i, almanac)

		if isSeedInRange(sourceSeed, seedRanges) {
			return i
		}
	}
	return -1
}
