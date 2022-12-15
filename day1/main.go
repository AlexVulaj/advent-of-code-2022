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
	max, current := 0, 0
	for _, line := range strings.Split(input, "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			if current > max {
				max = current
			}
			current = 0
		} else {
			current += num
		}
	}

	if current > max {
		max = current
	}

	return max
}

func p2(input string) int {
	first, second, third, current := 0, 0, 0, 0
	for _, line := range strings.Split(input, "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			if current > first {
				third = second
				second = first
				first = current
			} else if current > second {
				third = second
				second = current
			} else if current > third {
				third = current
			}
			current = 0
		} else {
			current += num
		}
	}

	if current > first {
		third = second
		second = first
		first = current
	} else if current > second {
		third = second
		second = current
	} else if current > third {
		third = current
	}

	return first + second + third
}
