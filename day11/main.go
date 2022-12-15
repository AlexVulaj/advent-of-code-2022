package main

import (
	"advent-of-code-2022/util"
	_ "embed"
	"sort"
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

type monkey struct {
	items       []int
	operation   []string
	test        int
	true        int
	false       int
	inspections int
}

func p1(input string) int {
	return solve(input, 20, 3)
}

func p2(input string) int {
	return solve(input, 10000, 1)
}

func solve(input string, rounds, inspectionRelief int) int {
	monkeys := parseMonkeys(input)

	manageWorry := 1
	for _, monkey := range monkeys {
		manageWorry *= monkey.test
	}

	for round := 0; round < rounds; round++ {
		for i := range monkeys {
			for j := range monkeys[i].items {
				monkeys[i].items[j] = inspect(monkeys[i].items[j], monkeys[i].operation)
				monkeys[i].inspections++
				monkeys[i].items[j] /= inspectionRelief
				monkeys[i].items[j] = monkeys[i].items[j] % manageWorry
				if checkTest(monkeys[i].items[j], monkeys[i].test) {
					monkeys[monkeys[i].true].items = append(monkeys[monkeys[i].true].items, monkeys[i].items[j])
				} else {
					monkeys[monkeys[i].false].items = append(monkeys[monkeys[i].false].items, monkeys[i].items[j])
				}
			}
			monkeys[i].items = []int{}
		}
	}
	var inspections []int
	for _, currentMonkey := range monkeys {
		inspections = append(inspections, currentMonkey.inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]
}

func inspect(old int, operation []string) int {
	operator := operation[0]
	var secondNum int
	if operation[1] == "old" {
		secondNum = old
	} else {
		secondNum, _ = strconv.Atoi(operation[1])
	}

	if operator == "+" {
		return old + secondNum
	}
	return old * secondNum
}

func checkTest(item int, test int) bool {
	return item%test == 0
}

func parseMonkeys(input string) []monkey {
	var monkeys []monkey
	splitInput := strings.Split(input, "\n")
	for i := 0; i < len(splitInput); i += 7 {
		currentMonkey := monkey{}
		currentMonkey.items = parseStartingItems(splitInput[i+1])
		currentMonkey.operation = parseOperation(splitInput[i+2])
		currentMonkey.test = parseTest(splitInput[i+3])
		currentMonkey.true = parseTrue(splitInput[i+4])
		currentMonkey.false = parseFalse(splitInput[i+5])

		monkeys = append(monkeys, currentMonkey)
	}
	return monkeys
}

func parseStartingItems(input string) []int {
	var items []int
	for _, val := range strings.Split(strings.Split(input, ":")[1], ",") {
		itemNum, _ := strconv.Atoi(strings.TrimSpace(val))
		items = append(items, itemNum)
	}
	return items
}

func parseOperation(input string) []string {
	return strings.Split(input, " ")[6:8]
}

func parseTest(input string) int {
	value, _ := strconv.Atoi(strings.Split(input, " ")[5])
	return value
}

func parseTrue(input string) int {
	value, _ := strconv.Atoi(strings.Split(input, " ")[9])
	return value
}

func parseFalse(input string) int {
	value, _ := strconv.Atoi(strings.Split(input, " ")[9])
	return value
}
