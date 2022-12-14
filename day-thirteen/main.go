package main

import (
	"advent-of-code-2022/util"
	_ "embed"
	"encoding/json"
	"reflect"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

const (
	rightOrderTrue  = 1
	rightOrderFalse = -1
	rightOrderSame  = 0
)

func main() {
	if util.ParsePartFlag() == 1 {
		util.PrintResult(p1(input))
	} else {
		util.PrintResult(p2(input))
	}
}

func p1(input string) int {
	var matchingIndices []int

	splitInput := strings.Split(input, "\n")
	currentPair := 0
	for lines := 0; lines < len(splitInput); lines += 3 {
		currentPair++
		var left []any
		_ = json.Unmarshal([]byte(splitInput[lines]), &left)
		var right []any
		_ = json.Unmarshal([]byte(splitInput[lines+1]), &right)

		if compareList(left, right) == rightOrderTrue {
			matchingIndices = append(matchingIndices, currentPair)
		}
	}

	sumIndices := 0
	for _, index := range matchingIndices {
		sumIndices += index
	}
	return sumIndices
}

func p2(input string) int {
	formattedInput := strings.Split(input, "\n")
	formattedInput = removeEmptyLines(formattedInput)
	formattedInput = append(formattedInput, "[[2]]", "[[6]]")

	var parsedInput [][]any
	for _, line := range formattedInput {
		var parsedLine []any
		_ = json.Unmarshal([]byte(line), &parsedLine)
		parsedInput = append(parsedInput, parsedLine)
	}

	sort.Slice(parsedInput, func(i, j int) bool { return compareList(parsedInput[i], parsedInput[j]) == rightOrderTrue })

	decoderIndices := make([]int, 0, 2)
	for i, line := range parsedInput {
		if reflect.DeepEqual(line, []any{[]any{2.0}}) || reflect.DeepEqual(line, []any{[]any{6.0}}) {
			decoderIndices = append(decoderIndices, i+1)
		}
	}
	return decoderIndices[0] * decoderIndices[1]
}

func removeEmptyLines(input []string) (parsedInput []string) {
	for _, line := range input {
		if line != "" {
			parsedInput = append(parsedInput, line)
		}
	}
	return
}

func compareList(left, right []any) int {
	leftLen, rightLen := len(left), len(right)
	var minLen int
	if leftLen < rightLen {
		minLen = leftLen
	} else {
		minLen = rightLen
	}

	for i := 0; i < minLen; i++ {
		var result int
		switch leftVal := left[i].(type) {
		case float64:
			switch rightVal := right[i].(type) {
			case float64:
				if leftVal < rightVal {
					result = rightOrderTrue
				}
				if leftVal > rightVal {
					result = rightOrderFalse
				}
				if leftVal == rightVal {
					result = rightOrderSame
				}
			case []any:
				result = compareList([]any{leftVal}, rightVal)
			}
		case []any:
			switch rightVal := right[i].(type) {
			case float64:
				result = compareList(leftVal, []any{rightVal})
			case []any:
				result = compareList(leftVal, rightVal)
			}
		}
		if result == rightOrderTrue {
			return rightOrderTrue
		}
		if result == rightOrderFalse {
			return rightOrderFalse
		}
	}
	if leftLen < rightLen {
		return rightOrderTrue
	}
	if leftLen > rightLen {
		return rightOrderFalse
	}
	return rightOrderSame
}
