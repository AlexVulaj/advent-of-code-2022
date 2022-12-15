package main

import (
	"advent-of-code-2022/util"
	_ "embed"
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
	points := 0
	for _, line := range strings.Split(input, "\n") {
		round := strings.Split(line, " ")
		opponent := round[0]
		me := round[1]

		if me == "X" {
			points += 1
			if opponent == "A" {
				points += 3
			} else if opponent == "C" {
				points += 6
			}
		} else if me == "Y" {
			points += 2
			if opponent == "A" {
				points += 6
			} else if opponent == "B" {
				points += 3
			}
		} else {
			points += 3
			if opponent == "B" {
				points += 6
			} else if opponent == "C" {
				points += 3
			}
		}
	}
	return points
}

func p2(input string) int {
	points := 0
	for _, line := range strings.Split(input, "\n") {
		round := strings.Split(line, " ")
		opponent := round[0]
		desiredResult := round[1]

		if desiredResult == "X" {
			if opponent == "A" {
				points += 3
			} else if opponent == "B" {
				points += 1
			} else {
				points += 2
			}
		} else if desiredResult == "Y" {
			points += 3
			if opponent == "A" {
				points += 1
			} else if opponent == "B" {
				points += 2
			} else {
				points += 3
			}
		} else {
			points += 6
			if opponent == "A" {
				points += 2
			} else if opponent == "B" {
				points += 3
			} else {
				points += 1
			}
		}
	}
	return points
}
