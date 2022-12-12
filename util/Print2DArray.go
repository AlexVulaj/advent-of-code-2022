package util

import "fmt"

func Print2dArray[K any](array [][]K) {
	for i, row := range array {
		for j, _ := range row {
			fmt.Print(array[i][j])
		}
		fmt.Println()
	}
}
