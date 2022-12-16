package main

import (
	"advent-of-code-2022/util"
	_ "embed"
	"fmt"
	set "github.com/AlexVulaj/go-set"
	"strings"
)

//go:embed input.txt
var input string

type Sensor struct {
	position      util.Point
	closestBeacon util.Point
}

type Line = [2]util.Point

func (s *Sensor) takenInRow(y int) set.Set[util.Point] {
	occupied := set.NewSet[util.Point]()

	diffY := y - s.position.Y
	for dx := -s.distance() + util.Abs(diffY); dx <= s.distance()-util.Abs(diffY); dx++ {
		occupied.Add(util.Point{X: s.position.X + dx, Y: y})
	}
	return occupied
}

func (s *Sensor) distance() int {
	return util.Abs(s.position.X-s.closestBeacon.X) + util.Abs(s.position.Y-s.closestBeacon.Y)
}

func (s *Sensor) boundaries() (topLeft, topRight, bottomRight, bottomLeft Line) {
	distance := s.distance() + 1

	left := util.Point{X: s.position.X - distance, Y: s.position.Y}
	top := util.Point{X: s.position.X, Y: s.position.Y - distance}
	right := util.Point{X: s.position.X + distance, Y: s.position.Y}
	bottom := util.Point{X: s.position.X, Y: s.position.Y + distance}

	topLeft[0] = left
	topLeft[1] = top

	topRight[0] = top
	topRight[1] = right

	bottomRight[0] = bottom
	bottomRight[1] = right

	bottomLeft[0] = left
	bottomLeft[1] = bottom

	return
}

func main() {
	if util.ParsePartFlag() == 1 {
		util.PrintResult(p1(input))
	} else {
		util.PrintResult(p2(input))
	}
}

func p1(input string) int {
	points := set.NewSet[util.Point]()
	sensors := parseLines(input)
	y := 2000000

	for _, sensor := range sensors {
		takenInRow := sensor.takenInRow(y)
		points.AddAll(takenInRow.ToSlice()...)
	}

	for _, sensor := range sensors {
		if sensor.closestBeacon.Y == y {
			points.Remove(sensor.closestBeacon)
		}
	}
	return points.Size()
}

func p2(input string) int {
	sensors := parseLines(input)
	y := 4000000

	for point := range calcIntersectionPoints(sensors, y) {
		invalid := false

		for _, sensor := range sensors {
			if util.Abs(sensor.position.X-point.X)+util.Abs(sensor.position.Y-point.Y) <= sensor.distance() {
				invalid = true
				break
			}
		}
		if !invalid {
			return point.X*4000000 + point.Y
		}
	}

	return 0
}

func calcIntersectionPoints(sensors []Sensor, bounds int) set.Set[util.Point] {
	points := set.NewSet(util.Point{}, util.Point{X: bounds}, util.Point{Y: bounds}, util.Point{X: bounds, Y: bounds})

	var ascLines, descLines []Line

	for _, sensor := range sensors {
		topLeft, topRight, bottomRight, bottomLeft := sensor.boundaries()

		ascLines = append(ascLines, topLeft, bottomRight)
		descLines = append(descLines, topRight, bottomLeft)
	}

	for _, ascLine := range ascLines {
		for _, descLine := range descLines {
			aIntersection := ascLine[0].Y - ascLine[0].X
			bIntersection := descLine[0].Y + descLine[0].X
			point := util.Point{X: (bIntersection - aIntersection) / 2, Y: (bIntersection + aIntersection) / 2}

			if point.X >= 0 && point.X <= bounds && point.Y >= 0 && point.Y <= bounds {
				points.Add(point)
			}
		}
	}
	return points
}

func parseLines(input string) []Sensor {
	lines := strings.Split(input, "\n")
	sensors := make([]Sensor, len(lines))

	for i, line := range lines {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)

		sensors[i] = Sensor{
			position:      util.Point{X: sensorX, Y: sensorY},
			closestBeacon: util.Point{X: beaconX, Y: beaconY},
		}
	}
	return sensors
}
