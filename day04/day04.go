package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day04(part string) {
	switch part {
	case "1":
		res := part1("./day04/input.txt")
		fmt.Println("The result for day 4 part 1 is", res)
	case "2":
		res := part2("./day04/input.txt")
		fmt.Println("The result for day 4 part 2 is", res)
	default:
		log.Fatal("Must run either part 1 or part 2")
	}
}

func part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	a := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), "")
		a = append(a, c)
	}

	return findMatches(a)
}

func part2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	a := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), "")
		a = append(a, c)
	}

	return findMatches2(a)
}

func findMatches(s [][]string) int {
	res := 0
	for i, x := range s {
		for j, _ := range x {
			if s[i][j] == "X" {
				// left to right
				if j < len(s[i])-3 && s[i][j+1] == "M" && s[i][j+2] == "A" && s[i][j+3] == "S" {
					res += 1
				}
				// right to left
				if j > 2 && s[i][j-1] == "M" && s[i][j-2] == "A" && s[i][j-3] == "S" {
					res += 1
				}
				// up to down
				if i < len(s)-3 && s[i+1][j] == "M" && s[i+2][j] == "A" && s[i+3][j] == "S" {
					res += 1
				}
				// down to up
				if i > 2 && s[i-1][j] == "M" && s[i-2][j] == "A" && s[i-3][j] == "S" {
					res += 1
				}
				// up-left to down-right
				if i < len(s)-3 && j < len(s[i])-3 && s[i+1][j+1] == "M" && s[i+2][j+2] == "A" && s[i+3][j+3] == "S" {
					res += 1
				}
				// up-right to down-left
				if i < len(s)-3 && j > 2 && s[i+1][j-1] == "M" && s[i+2][j-2] == "A" && s[i+3][j-3] == "S" {
					res += 1
				}
				// down-left to up-right
				if i > 2 && j < len(s[i])-3 && s[i-1][j+1] == "M" && s[i-2][j+2] == "A" && s[i-3][j+3] == "S" {
					res += 1
				}
				// down-right to up-left
				if i > 2 && j > 2 && s[i-1][j-1] == "M" && s[i-2][j-2] == "A" && s[i-3][j-3] == "S" {
					res += 1
				}
			}
		}
	}

	return res
}

func findMatches2(s [][]string) int {
	res := 0

	for i, x := range s {
		for j, _ := range x {
			if s[i][j] == "M" {
				// "M-S" top
				if i < len(s)-2 && j < len(s[i])-2 &&
					s[i+1][j+1] == "A" && s[i+2][j+2] == "S" && s[i+2][j] == "M" && s[i][j+2] == "S" {
					res += 1
				}
				// "M-M" top
				if i < len(s)-2 && j < len(s[i])-2 &&
					s[i+1][j+1] == "A" && s[i+2][j+2] == "S" && s[i+2][j] == "S" && s[i][j+2] == "M" {
					res += 1
				}
			}
			if s[i][j] == "S" {
				// "S-M" top
				if i < len(s)-2 && j < len(s[i])-2 &&
					s[i+1][j+1] == "A" && s[i+2][j+2] == "M" && s[i+2][j] == "S" && s[i][j+2] == "M" {
					res += 1
				}
				// "S-S" top
				if i < len(s)-2 && j < len(s[i])-2 &&
					s[i+1][j+1] == "A" && s[i+2][j+2] == "M" && s[i+2][j] == "M" && s[i][j+2] == "S" {
					res += 1
				}
			}
		}
	}

	return res
}
