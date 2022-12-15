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
	forest := parseInput(input)
	visible := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			if isVisible(forest, i, j) {
				visible++
			}
		}
	}
	return visible
}

func p2(input string) int {
	forest := parseInput(input)
	maxScore := 0
	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			currentScore := calculateScenicScore(forest, i, j)
			if currentScore > maxScore {
				maxScore = currentScore
			}
		}
	}
	return maxScore
}

func calculateScenicScore(forest [][]int, row, col int) int {
	upDog := 0
	for i := row - 1; i >= 0; i-- {
		upDog++
		if forest[i][col] >= forest[row][col] {
			break
		}
	}

	downDog := 0
	for i := row + 1; i < len(forest); i++ {
		downDog++
		if forest[i][col] >= forest[row][col] {
			break
		}
	}

	leftDog := 0
	for i := col - 1; i >= 0; i-- {
		leftDog++
		if forest[row][i] >= forest[row][col] {
			break
		}
	}

	rightDog := 0
	for i := col + 1; i < len(forest[row]); i++ {
		rightDog++
		if forest[row][i] >= forest[row][col] {
			break
		}
	}

	return upDog * downDog * leftDog * rightDog
}

func isVisible(forest [][]int, row, col int) bool {
	if row == 0 || col == 0 || row == len(forest)-1 || col == len(forest[row])-1 {
		return true
	}

	visible := true
	for i := row - 1; i >= 0; i-- {
		if forest[i][col] >= forest[row][col] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for i := row + 1; i < len(forest); i++ {
		if forest[i][col] >= forest[row][col] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for i := col - 1; i >= 0; i-- {
		if forest[row][i] >= forest[row][col] {
			visible = false
			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for i := col + 1; i < len(forest[row]); i++ {
		if forest[row][i] >= forest[row][col] {
			visible = false
			break
		}
	}

	return visible
}

func parseInput(input string) [][]int {
	splitInput := strings.Split(input, "\n")
	rowLen := len(splitInput[0])
	parsed := make([][]int, len(splitInput))
	for i := 0; i < len(parsed); i++ {
		parsed[i] = make([]int, rowLen)
	}

	for i := 0; i < len(parsed); i++ {
		for j := 0; j < rowLen; j++ {
			height, _ := strconv.Atoi(string(splitInput[i][j]))
			parsed[i][j] = height
		}
	}

	return parsed
}
