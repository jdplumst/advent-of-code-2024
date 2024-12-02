package main

import (
	"log"
	"os"

	. "github.com/jdplumst/advent-of-code-2024/day01"
	. "github.com/jdplumst/advent-of-code-2024/day02"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("Must run application with following format: go run . [day] [part]")
	}
	day := os.Args[1]
	part := os.Args[2]
	switch day {
	case "1":
		Day01(part)
	case "2":
		Day02(part)
	default:
		log.Fatal("Must run a valid day")
	}
}
