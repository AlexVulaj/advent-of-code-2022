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

func p1(input string) int {
	pairs := 0
	for _, line := range strings.Split(input, "\n") {
		if util.RangeFullyContains(parseLine(line)) {
			pairs++
		}
	}
	return pairs
}

func p2(input string) int {
	overlaps := 0
	for _, line := range strings.Split(input, "\n") {
		firstMin, firstMax, secondMin, secondMax := parseLine(line)

		if util.NumInRange(firstMin, secondMin, secondMax) || util.NumInRange(firstMax, secondMin, secondMax) || util.NumInRange(secondMin, firstMin, firstMax) || util.NumInRange(secondMax, firstMin, firstMax) {
			overlaps++
		}
	}
	return overlaps
}

func parseLine(line string) (int, int, int, int) {
	elves := strings.Split(line, ",")
	first := strings.Split(elves[0], "-")
	second := strings.Split(elves[1], "-")

	firstMin, _ := strconv.Atoi(first[0])
	firstMax, _ := strconv.Atoi(first[1])
	secondMin, _ := strconv.Atoi(second[0])
	secondMax, _ := strconv.Atoi(second[1])

	return firstMin, firstMax, secondMin, secondMax
}
