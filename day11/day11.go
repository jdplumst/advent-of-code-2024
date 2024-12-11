package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day11(part string) {
	switch part {
	case "1":
		res := part1("./day11/input.txt")
		fmt.Println("The result for day 11 part 1 is", res)
	case "2":
		res := part2("./day11/input.txt")
		fmt.Println("The result for day 11 part 2 is", res)
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

	rocks := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rocks = strings.Split(scanner.Text(), " ")
	}

	for count := 0; count < 25; count++ {
		rocks = blink(rocks)
	}

	return len(rocks)
}

func part2(input string) int {
	return 1
}

func blink(rocks []string) []string {
	newRocks := make([]string, 0)
	for i := range rocks {
		if rocks[i] == "0" {
			newRocks = append(newRocks, "1")
		} else if len(rocks[i])%2 == 0 {
			newRocks = append(newRocks, rocks[i][:len(rocks[i])/2])

			right, err := strconv.Atoi(rocks[i][len(rocks[i])/2:])
			if err != nil {
				log.Fatal(err)
			}
			newRocks = append(newRocks, strconv.Itoa(right))
		} else {
			num, err := strconv.Atoi(rocks[i])
			if err != nil {
				log.Fatal(err)
			}
			num *= 2024
			newRocks = append(newRocks, strconv.Itoa(num))
		}
	}

	return newRocks
}
