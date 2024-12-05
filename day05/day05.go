package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Day05(part string) {
	switch part {
	case "1":
		res := part1("./day05/input.txt")
		fmt.Println("The result for day 5 part 1 is", res)
	case "2":
		res := part2("./day05/input.txt")
		fmt.Println("The result for day 5 part 2 is", res)
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

	rulesCheck := true

	rules := make([][]int, 0)
	correctUpdates := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			rulesCheck = false
			continue
		}

		if rulesCheck {
			c := strings.Split(scanner.Text(), "|")
			left, err := strconv.Atoi(c[0])
			if err != nil {
				log.Fatal(err)
			}
			right, err := strconv.Atoi(c[1])
			if err != nil {
				log.Fatal(err)
			}
			x := []int{left, right}
			rules = append(rules, x)
		} else {
			c := strings.Split(scanner.Text(), ",")
			x := make([]int, 0)
			for _, e := range c {
				n, err := strconv.Atoi(e)
				if err != nil {
					log.Fatal(err)
				}
				x = append(x, n)
			}
			correct := checkCorrect(x, rules)
			if correct {
				correctUpdates = append(correctUpdates, x)
			}
		}
	}

	for _, e := range correctUpdates {
		res += e[len(e)/2]
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

	rulesCheck := true

	rules := make([][]int, 0)
	incorrectUpdates := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			rulesCheck = false
			continue
		}

		if rulesCheck {
			c := strings.Split(scanner.Text(), "|")
			left, err := strconv.Atoi(c[0])
			if err != nil {
				log.Fatal(err)
			}
			right, err := strconv.Atoi(c[1])
			if err != nil {
				log.Fatal(err)
			}
			x := []int{left, right}
			rules = append(rules, x)
		} else {
			c := strings.Split(scanner.Text(), ",")
			x := make([]int, 0)
			for _, e := range c {
				n, err := strconv.Atoi(e)
				if err != nil {
					log.Fatal(err)
				}
				x = append(x, n)
			}
			correct := checkCorrect(x, rules)
			if correct {
				continue
			}
			for {
				if correct {
					break
				}
				x = fixLine(x, rules)
				correct = checkCorrect(x, rules)
			}
			incorrectUpdates = append(incorrectUpdates, x)
		}
	}

	for _, e := range incorrectUpdates {
		res += e[len(e)/2]
	}

	return res
}

func checkCorrect(line []int, rules [][]int) bool {
	correct := true
	for i := 0; i < len(line); i++ {
		for j := i + 1; j < len(line); j++ {
			for k, _ := range rules {
				if rules[k][0] == line[j] && rules[k][1] == line[i] {
					correct = false
				}
			}
		}
	}

	return correct
}

func fixLine(line []int, rules [][]int) []int {
	for i := 0; i < len(line); i++ {
		for j := i + 1; j < len(line); j++ {
			for k, _ := range rules {
				if rules[k][0] == line[j] && rules[k][1] == line[i] {
					reflect.Swapper(line)(i, j)
				}
			}
		}
	}
	return line
}
