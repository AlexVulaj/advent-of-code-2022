package main

import (
	"advent-of-code-2022/util"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	if util.ParsePartFlag() == 1 {
		util.PrintResult(p1(input))
	} else {
		util.PrintResult(p2(input))
	}
}

func p1(input string) string {
	lines, stacks := parseLinesAndStacks(input)

	for i := 10; i < len(lines); i++ {
		parsedLine := strings.Split(lines[i], " ")
		moveCount, _ := strconv.Atoi(parsedLine[1])
		from := parsedLine[3]
		to := parsedLine[5]

		fromPointer := assignStack(from, stacks)
		toPointer := assignStack(to, stacks)

		for i := 0; i < moveCount; i++ {
			*toPointer = string((*fromPointer)[0]) + *toPointer
			*fromPointer = (*fromPointer)[1:]
		}
	}

	return parseReturn(stacks)
}

func p2(input string) string {
	lines, stacks := parseLinesAndStacks(input)

	for i := 10; i < len(lines); i++ {
		parsedLine := strings.Split(lines[i], " ")
		moveCount, _ := strconv.Atoi(parsedLine[1])
		from := parsedLine[3]
		to := parsedLine[5]

		fromPointer := assignStack(from, stacks)
		toPointer := assignStack(to, stacks)

		*toPointer = ((*fromPointer)[:moveCount]) + *toPointer
		*fromPointer = (*fromPointer)[moveCount:]
	}

	return parseReturn(stacks)
}

func parseLinesAndStacks(input string) ([]string, []string) {
	lines := strings.Split(input, "\n")
	stacks := make([]string, 9)
	for i := 0; i < 8; i++ {
		line := lines[i]

		for j := 0; j < 9; j++ {
			stacks[j] += string(line[j*4+1])
			stacks[j] = strings.TrimSpace(stacks[j])
		}
	}

	return lines, stacks
}

func assignStack(stack string, stacks []string) *string {
	stackInt, _ := strconv.Atoi(stack)
	return &stacks[stackInt-1]
}

func parseReturn(stacks []string) string {
	var result string
	for i := 0; i < 9; i++ {
		result += string(stacks[i][0])
	}
	return result
}
