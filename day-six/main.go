package main

import (
	"advent-of-code-2022/util"
	_ "embed"
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
	return findMarker(input, 4)
}

func p2(input string) int {
	return findMarker(input, 14)
}

func findMarker(input string, sequenceLength int) int {
	for i := sequenceLength - 1; i < len(input); i++ {
		strings := make(map[uint8]bool)
		for j := i - sequenceLength + 1; j <= i; j++ {
			if _, ok := strings[input[j]]; !ok {
				strings[input[j]] = true
			} else {
				break
			}
		}
		if len(strings) == sequenceLength {
			return i + 1
		}
	}
	return -1
}
