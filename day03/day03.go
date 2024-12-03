package day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day03(part string) {
	switch part {
	case "1":
		res := part1("./day03/input.txt")
		fmt.Println("The result for day 3 part 1 is", res)
	case "2":
		res := part2("./day03/input.txt")
		fmt.Println("The result for day 3 part 2 is", res)
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
	check := ""
	a := 0
	b := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		if scanner.Text() == "m" {
			check = "m"
		} else if scanner.Text() == "u" && check == "m" {
			check = "mu"
		} else if scanner.Text() == "l" && check == "mu" {
			check = "mul"
		} else if scanner.Text() == "(" && check == "mul" {
			check = "mul("
		} else if digitCheck.MatchString(scanner.Text()) && check == "mul(" {
			digit, _ := strconv.Atoi(scanner.Text())
			a = a * 10
			a += digit
		} else if digitCheck.MatchString(scanner.Text()) && check == "mul(," {
			digit, _ := strconv.Atoi(scanner.Text())
			b = b * 10
			b += digit
		} else if scanner.Text() == "," && check == "mul(" {
			check = "mul(,"
		} else if scanner.Text() == ")" && check == "mul(," {
			check = ""
			res += a * b
			a = 0
			b = 0
		} else {
			check = ""
			a = 0
			b = 0
		}

	}
	return res
}

func part2(input string) int { return 1 }

var digitCheck = regexp.MustCompile(`^[0-9]+$`)
