package main

import (
	"advent-of-code-2022/util"
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input.txt
	input  string
	source = util.Point{X: 500, Y: 0}
)

func main() {
	if util.ParsePartFlag() == 1 {
		util.PrintResult(p1(input))
	} else {
		util.PrintResult(p2(input))
	}
}

type Line = []util.Point

func p1(input string) int {
	return process(getPointsForLines(parseInput(input)), false)
}

func p2(input string) int {
	return process(getPointsForLines(parseInput(input)), true)
}

func process(rocks map[util.Point]bool, shouldFloor bool) int {
	rocksAndSand := map[util.Point]bool{}
	for rock := range rocks {
		rocksAndSand[rock] = true
	}

	lowest := findLowest(rocksAndSand)
	floor := lowest + 2

	settled := 0
	control := false

	for !control {
		sand := source
		hasSettled := false
		for !hasSettled {
			if !shouldFloor && sand.Y > lowest {
				control = true
				break
			}

			if shouldFloor && sand.Y == floor-1 {
				hasSettled = true
				rocksAndSand[sand] = true
				settled++
			}

			if _, ok := rocksAndSand[util.Point{X: sand.X, Y: sand.Y + 1}]; !ok {
				sand = util.Point{X: sand.X, Y: sand.Y + 1}
				continue
			}
			if _, ok := rocksAndSand[util.Point{X: sand.X - 1, Y: sand.Y + 1}]; !ok {
				sand = util.Point{X: sand.X - 1, Y: sand.Y + 1}
				continue
			}
			if _, ok := rocksAndSand[util.Point{X: sand.X + 1, Y: sand.Y + 1}]; !ok {
				sand = util.Point{X: sand.X + 1, Y: sand.Y + 1}
			} else {
				hasSettled = true
				rocksAndSand[sand] = true
				settled++
			}
		}

		if hasSettled && sand == source {
			control = true
		}
	}
	return settled
}

func findLowest(rocks map[util.Point]bool) int {
	lowest := 0

	for rock := range rocks {
		if rock.Y > lowest {
			lowest = rock.Y
		}
	}
	return lowest
}

func parseInput(input string) []Line {
	parsedLines := strings.Split(input, "\n")

	lines := make([]Line, len(parsedLines))

	for i, line := range parsedLines {
		points := strings.Split(line, " -> ")
		lines[i] = make(Line, len(points))
		for j, point := range points {
			lines[i][j] = parsePoint(point)
		}
	}
	return lines
}

func parsePoint(point string) util.Point {
	var x, y int
	fmt.Sscanf(point, "%d,%d", &x, &y)
	return util.Point{X: x, Y: y}
}

func getPointsForLines(lines []Line) map[util.Point]bool {
	points := map[util.Point]bool{}
	for _, line := range lines {
		for point := range getPointsForLine(line) {
			points[point] = true
		}
	}
	return points
}

func getPointsForLine(line Line) map[util.Point]bool {
	points := map[util.Point]bool{}
	for i := 0; i < len(line)-1; i++ {
		for point := range getPointsBetween(line[i], line[i+1]) {
			points[point] = true
		}
	}
	return points
}

func getPointsBetween(a, b util.Point) map[util.Point]bool {
	points := map[util.Point]bool{a: true, b: true}

	if a.X == b.X && a.Y > b.Y {
		for y := a.Y - 1; y > b.Y; y-- {
			points[util.Point{X: a.X, Y: y}] = true
		}
	}
	if a.X == b.X && a.Y < b.Y {
		for y := a.Y + 1; y < b.Y; y++ {
			points[util.Point{X: a.X, Y: y}] = true
		}
	}
	if a.X > b.X && a.Y == b.Y {
		for x := a.X - 1; x > b.X; x-- {
			points[util.Point{X: x, Y: a.Y}] = true
		}
	}
	if a.X < b.X && a.Y == b.Y {
		for x := a.X + 1; x < b.X; x++ {
			points[util.Point{X: x, Y: a.Y}] = true
		}
	}

	return points
}
