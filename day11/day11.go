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

type RockBlinks struct {
	Rock   string
	Blinks int
}

func part2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := 0

	rocks := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rocks = strings.Split(scanner.Text(), " ")
	}

	memo := map[RockBlinks]int{}

	for i := range rocks {
		res += blinkWithRecursion(rocks[i], 75, memo)
	}

	return res
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

func blinkWithRecursion(rock string, blinks int, memo map[RockBlinks]int) int {
	if blinks == 0 {
		return 1
	}

	cached := memo[RockBlinks{Rock: rock, Blinks: blinks}]
	if cached != 0 {
		return cached
	}

	sum := 0

	if rock == "0" {
		sum += blinkWithRecursion("1", blinks-1, memo)
	} else if len(rock)%2 == 0 {
		right, err := strconv.Atoi(rock[len(rock)/2:])
		if err != nil {
			log.Fatal(err)
		}
		sum += blinkWithRecursion(rock[:len(rock)/2], blinks-1, memo) + blinkWithRecursion(strconv.Itoa(right), blinks-1, memo)
	} else {
		num, err := strconv.Atoi(rock)
		if err != nil {
			log.Fatal(err)
		}
		num *= 2024
		sum += blinkWithRecursion(strconv.Itoa(num), blinks-1, memo)
	}

	memo[RockBlinks{Rock: rock, Blinks: blinks}] = sum

	return sum
}
