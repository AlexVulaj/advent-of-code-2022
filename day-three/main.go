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

var priorityValues = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52}

func p1(input string) int {
	priorities := 0
	for _, line := range strings.Split(input, "\n") {
		priorities += findLinePriority(line)
	}
	return priorities
}

func findLinePriority(line string) int {
	lineLength := len(line)
	firstHalf := line[:lineLength/2]
	secondHalf := line[lineLength/2:]
	for _, firstChar := range strings.Split(firstHalf, "") {
		for _, secondChar := range strings.Split(secondHalf, "") {
			if firstChar == secondChar {
				return priorityValues[firstChar]
			}
		}
	}
	return 0
}

func p2(input string) int {
	priorities := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 3 {
		priorities += findGroupPriority(lines[i], lines[i+1], lines[i+2])
	}
	return priorities
}

func findGroupPriority(first, second, third string) int {
	for _, firstChar := range strings.Split(first, "") {
		for _, secondChar := range strings.Split(second, "") {
			if firstChar == secondChar {
				for _, thirdChar := range strings.Split(third, "") {
					if secondChar == thirdChar {
						return priorityValues[firstChar]
					}
				}
			}
		}
	}
	return 0
}
