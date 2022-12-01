package util

import (
	"flag"
	"log"
)

func ParsePartFlag() int {
	var part int
	flag.IntVar(&part, "part", 0, "part 1 or 2")
	flag.Parse()
	if part != 1 && part != 2 {
		log.Fatal("'--part' flag must be provided with a value of 1 or 2")
	}
	return part
}
