package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day10(part string) {
	switch part {
	case "1":
		res := part1("./day10/input.txt")
		fmt.Println("The result for day 10 part 1 is", res)
	case "2":
		res := part2("./day10/input.txt")
		fmt.Println("The result for day 10 part 2 is", res)
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

	topMap := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), "")
		topMap = append(topMap, x)
	}
	tracked := make([][]bool, len(topMap))
	for i := range topMap {
		tracked[i] = make([]bool, len(topMap[i]))
	}

	for i := 0; i < len(topMap); i++ {
		for j := 0; j < len(topMap[i]); j++ {
			if topMap[i][j] == "0" {
				res += getScore(topMap, tracked, 0, i, j)
				tracked = resetTracked(tracked)
			}
		}
	}

	return res
}

func part2(input string) int {
	return 1
}

func getScore(topMap [][]string, tracked [][]bool, scale, i, j int) int {
	score := 0

	if scale == 9 && !tracked[i][j] {
		tracked[i][j] = true
		return 1
	}

	if i > 0 {
		top, err := strconv.Atoi(topMap[i-1][j])
		if err != nil {
			log.Fatal(err)
		}
		if top == scale+1 {
			score += getScore(topMap, tracked, top, i-1, j)
		}
	}
	if i < len(topMap)-1 {
		bot, err := strconv.Atoi(topMap[i+1][j])
		if err != nil {
			log.Fatal(err)
		}
		if bot == scale+1 {
			score += getScore(topMap, tracked, bot, i+1, j)
		}
	}
	if j > 0 {
		left, err := strconv.Atoi(topMap[i][j-1])
		if err != nil {
			log.Fatal(err)
		}
		if left == scale+1 {
			score += getScore(topMap, tracked, left, i, j-1)
		}
	}
	if j < len(topMap[i])-1 {
		right, err := strconv.Atoi(topMap[i][j+1])
		if err != nil {
			log.Fatal(err)
		}
		if right == scale+1 {
			score += getScore(topMap, tracked, right, i, j+1)
		}
	}

	return score
}

func resetTracked(tracked [][]bool) [][]bool {
	for i := 0; i < len(tracked); i++ {
		for j := 0; j < len(tracked[i]); j++ {
			tracked[i][j] = false
		}
	}

	return tracked
}
