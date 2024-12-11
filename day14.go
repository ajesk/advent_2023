package main

type rockStack []rune

func (s *rockStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *rockStack) Size() int {
	return len(*s)
}

func (s *rockStack) Push(ch rune) {
	*s = append(*s, ch) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *rockStack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return rune(0), false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *rockStack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return rune(0), false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

func day14Part1() int {
	fileContents, err := ReadFileLines("./data/day14.txt")
	print(fileContents)

	if err != nil {
		return 0
	}

	return rollAndCount(fileContents)
}

func rollAndCount(fileContents []string) int {
	startingIndex := len(fileContents)
	rolledRocks := parseRocks(fileContents)

	return tallyRocks(rolledRocks, startingIndex)
}

func day14Part2() int {
	fileContents, err := ReadFileLines("./data/day14.txt")
	print(fileContents)

	if err != nil {
		return 0
	}

	return 0
}

func tallyRocks(cols []*rockStack, startingIndex int) int {
	count := 0
	for _, s := range cols {

		for !s.IsEmpty() {
			currentRock, _ := s.Pop()
			currentRowVal := startingIndex - s.Size()

			if currentRock == 'O' {
				count += currentRowVal
			}
		}
	}
	return count
}

func parseRocks(lines []string) []*rockStack {
	cols := []*rockStack{}
	for i := 0; i < len(lines[0]); i++ {
		cols = append(cols, &rockStack{})
	}

	for _, line := range lines {
		for i, ch := range line {
			s := cols[i]
			removedDots := 0
			if ch == 'O' {
				nextCh, _ := s.Peek()
				for !s.IsEmpty() && nextCh == '.' {
					s.Pop()
					removedDots++
					nextCh, _ = s.Peek()
				}

			}
			s.Push(ch)
			for removedDots != 0 {
				s.Push('.')
				removedDots--
			}
		}
	}

	return cols
}
