package main

import (
	"log"
	"os"

	. "github.com/jdplumst/advent-of-code-2024/day01"
	. "github.com/jdplumst/advent-of-code-2024/day02"
	. "github.com/jdplumst/advent-of-code-2024/day03"
	. "github.com/jdplumst/advent-of-code-2024/day04"
	. "github.com/jdplumst/advent-of-code-2024/day05"
	. "github.com/jdplumst/advent-of-code-2024/day06"
	. "github.com/jdplumst/advent-of-code-2024/day07"
	. "github.com/jdplumst/advent-of-code-2024/day08"
	. "github.com/jdplumst/advent-of-code-2024/day09"
	. "github.com/jdplumst/advent-of-code-2024/day10"
	. "github.com/jdplumst/advent-of-code-2024/day11"
	. "github.com/jdplumst/advent-of-code-2024/day12"
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
	case "3":
		Day03(part)
	case "4":
		Day04(part)
	case "5":
		Day05(part)
	case "6":
		Day06(part)
	case "7":
		Day07(part)
	case "8":
		Day08(part)
	case "9":
		Day09(part)
	case "10":
		Day10(part)
	case "11":
		Day11(part)
	case "12":
		Day12(part)
	default:
		log.Fatal("Must run a valid day")
	}
}
