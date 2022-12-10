package main

import (
	"advent-of-code-2022/util"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	if util.ParsePartFlag() == 1 {
		util.PrintResult(p1(input))
	} else {
		fmt.Println("Answer:")
		print2dArray(p2(input))
		fmt.Println()
	}
}

func p1(input string) int {
	importantCycles := map[int]int{20: 0, 60: 0, 100: 0, 140: 0, 180: 0, 220: 0}
	cycle := 1
	xRegister := 1

	for _, line := range strings.Split(input, "\n") {
		step := strings.Split(line, " ")
		if step[0] == "noop" {
			cycle++
			if _, ok := importantCycles[cycle]; ok {
				importantCycles[cycle] = xRegister
			}
		} else {
			cycle++
			if _, ok := importantCycles[cycle]; ok {
				importantCycles[cycle] = xRegister
			}
			cycle++
			xChange, _ := strconv.Atoi(step[1])
			xRegister += xChange
			if _, ok := importantCycles[cycle]; ok {
				importantCycles[cycle] = xRegister
			}
		}
	}

	sum := 0
	for key, value := range importantCycles {
		sum += key * value
	}
	return sum
}

func p2(input string) [][]string {
	lit := "#"
	dark := "."
	crt := make([][]string, 6)
	for i := range crt {
		crt[i] = make([]string, 40)
		for j := 0; j < 40; j++ {
			crt[i][j] = dark
		}
	}

	xRegister := 1
	cycle := 0
	positionX := cycle % 40
	positionY := cycle / 40

	for _, line := range strings.Split(input, "\n") {
		step := strings.Split(line, " ")
		if xRegister == positionX || xRegister == positionX-1 || xRegister == positionX+1 {
			crt[positionY][positionX] = lit
		}
		if step[0] == "noop" {
			cycle, positionX, positionY = incrementCycle(cycle)
			if xRegister == positionX || xRegister == positionX-1 || xRegister == positionX+1 {
				crt[positionY][positionX] = lit
			}
		} else {
			cycle, positionX, positionY = incrementCycle(cycle)
			if xRegister == positionX || xRegister == positionX-1 || xRegister == positionX+1 {
				crt[positionY][positionX] = lit
			}
			cycle, positionX, positionY = incrementCycle(cycle)
			xChange, _ := strconv.Atoi(step[1])
			xRegister += xChange
			if xRegister == positionX || xRegister == positionX-1 || xRegister == positionX+1 {
				crt[positionY][positionX] = lit
			}
		}
	}

	fmt.Println(crt)
	return crt
}

func print2dArray[K any](array [][]K) {
	for i, row := range array {
		for j, _ := range row {
			fmt.Print(array[i][j])
		}
		fmt.Println()
	}
}

func incrementCycle(currentCycle int) (int, int, int) {
	newCycle := currentCycle + 1
	return newCycle, newCycle % 40, newCycle / 40
}
