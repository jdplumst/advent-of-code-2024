package main

import (
	"os"

	. "github.com/jdplumst/advent-of-code-2024/day01"
)

func main() {
	day := os.Args[1]
	part := os.Args[2]
	switch day {
	case "1":
		Day01(part)
	default:
		panic("Must run a valid day")
	}
}
