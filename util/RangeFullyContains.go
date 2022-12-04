package util

func RangeFullyContains(firstMin, firstMax, secondMin, secondMax int) bool {
	return firstMin <= secondMin && firstMax >= secondMax || secondMin <= firstMin && secondMax >= firstMax
}
