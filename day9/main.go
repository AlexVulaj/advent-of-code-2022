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
	return solve(input, 2)
}

func p2(input string) int {
	return solve(input, 10)
}

func solve(input string, knotCount int) int {
	knots := make([]util.Point, knotCount)
	visited := map[string]struct{}{"0,0": {}}

	for _, line := range strings.Split(input, "\n") {
		step := strings.Split(line, " ")
		direction := step[0]
		moveCount, _ := strconv.Atoi(step[1])
		for i := 0; i < moveCount; i++ {
			knots[0].X, knots[0].Y = moveHead(knots[0], direction)
			for j := 0; j < len(knots)-1; j++ {
				if !isAdjacent(knots[j], knots[j+1]) {
					knots[j+1].X, knots[j+1].Y = moveTail(knots[j], knots[j+1])
					key := strconv.Itoa(knots[j+1].X) + "," + strconv.Itoa(knots[j+1].Y)
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

func isAdjacent(positionOne, positionTwo util.Point) bool {
	return util.Abs(positionOne.X-positionTwo.X) < 2 && util.Abs(positionOne.Y-positionTwo.Y) < 2
}

func moveHead(head util.Point, direction string) (int, int) {
	if direction == "U" {
		return head.X, head.Y + 1
	}
	if direction == "R" {
		return head.X + 1, head.Y
	}
	if direction == "D" {
		return head.X, head.Y - 1
	}
	if direction == "L" {
		return head.X - 1, head.Y
	}
	return head.X, head.Y
}

func moveTail(head, tail util.Point) (int, int) {
	if head.Y == tail.Y {
		if head.X > tail.X {
			return tail.X + 1, tail.Y
		}
		if head.X < tail.X {
			return tail.X - 1, tail.Y
		}
	}

	if head.X == tail.X {
		if head.Y > tail.Y {
			return tail.X, tail.Y + 1
		}
		if head.Y < tail.Y {
			return tail.X, tail.Y - 1
		}
	}

	if head.X > tail.X && head.Y > tail.Y {
		return tail.X + 1, tail.Y + 1
	}

	if head.X > tail.X && head.Y < tail.Y {
		return tail.X + 1, tail.Y - 1
	}

	if head.X < tail.X && head.Y < tail.Y {
		return tail.X - 1, tail.Y - 1
	}

	if head.X < tail.X && head.Y > tail.Y {
		return tail.X - 1, tail.Y + 1
	}

	return tail.X, tail.Y
}
