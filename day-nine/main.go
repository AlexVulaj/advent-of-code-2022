package main

import (
	"advent-of-code-2022/util"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type position struct {
	x int
	y int
}

func main() {
	if util.ParsePartFlag() == 1 {
		util.PrintResult(p1(input))
	} else {
		util.PrintResult(p2(input))
	}
}

func p1(input string) int {
	return solve(input, 2)
}

func p2(input string) int {
	return solve(input, 10)
}

func solve(input string, knotCount int) int {
	knots := make([]position, knotCount)
	visited := map[string]struct{}{"0,0": {}}

	for _, line := range strings.Split(input, "\n") {
		step := strings.Split(line, " ")
		direction := step[0]
		moveCount, _ := strconv.Atoi(step[1])
		for i := 0; i < moveCount; i++ {
			knots[0].x, knots[0].y = moveHead(knots[0], direction)
			for j := 0; j < len(knots)-1; j++ {
				if !isAdjacent(knots[j], knots[j+1]) {
					knots[j+1].x, knots[j+1].y = moveTail(knots[j], knots[j+1])
					key := strconv.Itoa(knots[j+1].x) + "," + strconv.Itoa(knots[j+1].y)
					if j+1 == knotCount-1 {
						_, ok := visited[key]
						if !ok {
							visited[key] = struct{}{}
						}
					}
				}
			}
		}
	}
	return len(visited)
}

func isAdjacent(positionOne, positionTwo position) bool {
	return util.Abs(positionOne.x-positionTwo.x) < 2 && util.Abs(positionOne.y-positionTwo.y) < 2
}

func moveHead(head position, direction string) (int, int) {
	if direction == "U" {
		return head.x, head.y + 1
	}
	if direction == "R" {
		return head.x + 1, head.y
	}
	if direction == "D" {
		return head.x, head.y - 1
	}
	if direction == "L" {
		return head.x - 1, head.y
	}
	return head.x, head.y
}

func moveTail(head, tail position) (int, int) {
	if head.y == tail.y {
		if head.x > tail.x {
			return tail.x + 1, tail.y
		}
		if head.x < tail.x {
			return tail.x - 1, tail.y
		}
	}

	if head.x == tail.x {
		if head.y > tail.y {
			return tail.x, tail.y + 1
		}
		if head.y < tail.y {
			return tail.x, tail.y - 1
		}
	}

	if head.x > tail.x && head.y > tail.y {
		return tail.x + 1, tail.y + 1
	}

	if head.x > tail.x && head.y < tail.y {
		return tail.x + 1, tail.y - 1
	}

	if head.x < tail.x && head.y < tail.y {
		return tail.x - 1, tail.y - 1
	}

	if head.x < tail.x && head.y > tail.y {
		return tail.x - 1, tail.y + 1
	}

	return tail.x, tail.y
}
