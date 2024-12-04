package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day02(part string) {
	switch part {
	case "1":
		res := part1("./day02/input.txt")
		fmt.Println("The result for day 2 part 1 is", res)
	case "2":
		res := part2("./day02/input.txt")
		fmt.Println("The result for day 2 part 2 is", res)
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

	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		line := strings.Split(scanner.Text(), " ")

		safe := true
		direction := "increasing"

		for idx, _ := range line {
			if idx > 0 {
				prev, err := strconv.Atoi(line[idx-1])
				if err != nil {
					log.Fatal(err)
				}
				curr, err := strconv.Atoi(line[idx])
				if err != nil {
					log.Fatal(err)
				}

				if idx == 1 {
					if curr > prev {
						direction = "increasing"
					} else if curr < prev {
						direction = "decreasing"
					} else {
						safe = false
					}
				}

				switch direction {
				case "increasing":
					if curr <= prev {
						safe = false
					}
				case "decreasing":
					if curr >= prev {
						safe = false
					}
				}

				if absoluteValue(prev, curr) > 3 {
					safe = false
				}
			}
		}

		if safe {
			res += 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
func part2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		line := strings.Split(scanner.Text(), " ")

		safe, _ := checkSafe(line)

		if safe {
			res += 1
		} else {
			for idx, _ := range line {
				newLine := append([]string{}, line...)
				newLine = append(newLine[:idx], newLine[idx+1:]...)
				safe, _ := checkSafe(newLine)
				if safe {
					res += 1
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

func absoluteValue(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = diff * -1
	}
	return diff
}

func checkSafe(line []string) (bool, []int) {
	safe := true
	unsafeIndexes := make([]int, 0)
	direction := "increasing"

	for idx, _ := range line {
		if idx > 0 {
			prev, err := strconv.Atoi(line[idx-1])
			if err != nil {
				log.Fatal(err)
			}
			curr, err := strconv.Atoi(line[idx])
			if err != nil {
				log.Fatal(err)
			}

			if idx == 1 {
				if curr > prev {
					direction = "increasing"
				} else if curr < prev {
					direction = "decreasing"
				} else {
					safe = false
					unsafeIndexes = append(unsafeIndexes, idx-1)
					unsafeIndexes = append(unsafeIndexes, idx)
				}
			}

			switch direction {
			case "increasing":
				if curr <= prev {
					safe = false
					unsafeIndexes = append(unsafeIndexes, idx-1)
					unsafeIndexes = append(unsafeIndexes, idx)
				}
			case "decreasing":
				if curr >= prev {
					safe = false
					unsafeIndexes = append(unsafeIndexes, idx-1)
					unsafeIndexes = append(unsafeIndexes, idx)
				}
			}

			if absoluteValue(prev, curr) > 3 {
				safe = false
				unsafeIndexes = append(unsafeIndexes, idx-1)
				unsafeIndexes = append(unsafeIndexes, idx)
			}
		}
	}

	return safe, unsafeIndexes
}
