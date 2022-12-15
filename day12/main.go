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

var diffs = [4][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func p1(input string) int {
	heightMap := parseInput(input)

	var queue [][3]int
	for i, row := range heightMap {
		for j, cell := range row {
			if cell == "S" {
				queue = append(queue, [3]int{i, j, 0})
				break
			}
		}
	}
	seen := map[[2]int]bool{}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if seen[[2]int{front[0], front[1]}] {
			continue
		}
		seen[[2]int{front[0], front[1]}] = true

		if heightMap[front[0]][front[1]] == "E" {
			return front[2]
		}
		for _, d := range diffs {
			nextR, nextC := front[0]+d[0], front[1]+d[1]
			if nextR >= 0 && nextR < len(heightMap) && nextC >= 0 && nextC < len(heightMap[0]) {
				letterDiff := heightDifference(heightMap[front[0]][front[1]], heightMap[nextR][nextC])

				if letterDiff <= 1 {
					queue = append(queue, [3]int{nextR, nextC, front[2] + 1})
				}
			}
		}
	}

	return -1
}

func p2(input string) int {
	heightMap := parseInput(input)

	var queue [][3]int
	for r, rows := range heightMap {
		for c, cell := range rows {
			if cell == "E" {
				queue = append(queue, [3]int{r, c, 0})
				break
			}
		}
	}
	seen := map[[2]int]bool{}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if seen[[2]int{front[0], front[1]}] {
			continue
		}
		seen[[2]int{front[0], front[1]}] = true

		if heightMap[front[0]][front[1]] == "a" {
			return front[2]
		}
		for _, d := range diffs {
			nextR, nextC := front[0]+d[0], front[1]+d[1]
			if nextR >= 0 && nextR < len(heightMap) && nextC >= 0 && nextC < len(heightMap[0]) {
				letterDiff := heightDifference(heightMap[front[0]][front[1]], heightMap[nextR][nextC])

				if letterDiff >= -1 {
					queue = append(queue, [3]int{nextR, nextC, front[2] + 1})
				}
			}
		}
	}

	return -1
}

func heightDifference(x, y string) int {
	if x == "S" {
		x = "a"
	}
	if y == "S" {
		y = "a"
	}
	if y == "E" {
		y = "z"
	}
	if x == "E" {
		x = "z"
	}
	return int(y[0]) - int(x[0])
}

func parseInput(input string) (heightMap [][]string) {
	for _, line := range strings.Split(input, "\n") {
		heightMap = append(heightMap, strings.Split(line, ""))
	}
	return heightMap
}
