package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day01(part string) {
	switch part {
	case "1":
		res := part1("./day01/input.txt")
		fmt.Println("The result for day 1 part 1 is", res)
	case "2":
		res := part2("./day01/input.txt")
		fmt.Println("The result for day 1 part 2 is", res)
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

	list1, list2 := make([]int, 0), make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		line := strings.Split(scanner.Text(), "   ")

		a, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}

		b, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, a)
		list2 = append(list2, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for idx, _ := range list1 {
		diff := list1[idx] - list2[idx]
		if diff < 0 {
			diff = diff * -1
		}
		res += diff
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

	list1, list2 := make([]int, 0), make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		line := strings.Split(scanner.Text(), "   ")

		a, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}

		b, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, a)
		list2 = append(list2, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for _, elem1 := range list1 {
		count := 0
		for _, elem2 := range list2 {
			if elem1 == elem2 {
				count += 1
			}
		}
		res += (elem1 * count)
	}

	return res
}
