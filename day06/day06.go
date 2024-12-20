package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day06(part string) {
	switch part {
	case "1":
		res := part1("./day06/input.txt")
		fmt.Println("The result for day 6 part 1 is", res)
	case "2":
		res := part2("./day06/input.txt")
		fmt.Println("The result for day 6 part 2 is", res)
	default:
		log.Fatal("Must run either part 1 or part 2")
	}
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

func part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	guard_map := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), "")
		guard_map = append(guard_map, c)
	}

	x, y, d, err := findGuard(guard_map)
	if err != nil {
		log.Fatal(err)
	}

	for {
		completed := false
		x, y, d, completed = markPath(guard_map, x, y, d)
		if completed {
			break
		}
	}

	return countPositions(guard_map)
}

func part2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := 0

	guard_map := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), "")
		guard_map = append(guard_map, c)
	}

	for i := 0; i < len(guard_map); i++ {
		for j := 0; j < len(guard_map); j++ {
			loop := checkLoop(guard_map, i, j)
			if loop {
				res++
			}
		}
	}

	return res
}

func findGuard(guard_map [][]string) (int, int, direction, error) {
	for i := 0; i < len(guard_map); i++ {
		for j := 0; j < len(guard_map[i]); j++ {
			if guard_map[i][j] == "^" {
				return i, j, direction(up), nil
			} else if guard_map[i][j] == ">" {
				return i, j, direction(right), nil
			} else if guard_map[i][j] == "v" {
				return i, j, direction(down), nil
			} else if guard_map[i][j] == "<" {
				return i, j, direction(left), nil
			}
		}
	}

	return -1, -1, -1, fmt.Errorf("could not find guard")
}

func markPath(guard_map [][]string, x, y int, d direction) (int, int, direction, bool) {
	for {
		switch d {
		case direction(up):
			if x < 0 {
				return -1, -1, -1, true
			}
			guard_map[x][y] = "X"
			x--
			if x >= 0 && guard_map[x][y] == "#" {
				guard_map[x+1][y] = ">"
				return x + 1, y, direction(right), false
			} else if x >= 0 {
				guard_map[x][y] = "^"
			}
		case direction(right):
			if y >= len(guard_map[x]) {
				return -1, -1, -1, true
			}
			guard_map[x][y] = "X"
			y++
			if y < len(guard_map[x]) && guard_map[x][y] == "#" {
				guard_map[x][y-1] = "v"
				return x, y - 1, direction(down), false
			} else if y < len(guard_map[x]) {
				guard_map[x][y] = ">"
			}
		case direction(down):
			if x >= len(guard_map) {
				return -1, -1, -1, true
			}
			guard_map[x][y] = "X"
			x++
			if x < len(guard_map) && guard_map[x][y] == "#" {
				guard_map[x-1][y] = "<"
				return x - 1, y, direction(left), false
			} else if x < len(guard_map) {
				guard_map[x][y] = "v"
			}
		case direction(left):
			if y < 0 {
				return -1, -1, -1, true
			}
			guard_map[x][y] = "X"
			y--
			if y >= 0 && guard_map[x][y] == "#" {
				guard_map[x][y+1] = "<"
				return x, y + 1, direction(up), false
			} else if y >= 0 {
				guard_map[x][y] = "^"
			}
		}
	}
}

func countPositions(guard_map [][]string) int {
	count := 0

	for i := 0; i < len(guard_map); i++ {
		for j := 0; j < len(guard_map[i]); j++ {
			if guard_map[i][j] == "X" {
				count++
			}
		}
	}

	return count
}

func checkLoop(guard_map [][]string, x, y int) bool {
	if guard_map[x][y] != "." {
		return false
	}

	gx, gy, d, err := findGuard(guard_map)
	if err != nil {
		log.Fatal(err)
	}

	obstructed_map := make([][]string, len(guard_map))
	for i := range guard_map {
		obstructed_map[i] = make([]string, len(guard_map[i]))
		copy(obstructed_map[i], guard_map[i])
	}

	obstructed_map[x][y] = "#"

	for {
		switch d {
		case direction(up):
			gx--
			if gx < 0 {
				return false
			}
			if strings.Contains(obstructed_map[gx][gy], "U") {
				return true
			} else {
				obstructed_map[gx][gy] += "U"
				if strings.Contains(obstructed_map[gx][gy], "#") {
					d = direction(right)
					gx++
				}
			}
		case direction(right):
			gy++
			if gy >= len(obstructed_map[gx]) {
				return false
			}
			if strings.Contains(obstructed_map[gx][gy], "R") {
				return true
			} else {
				obstructed_map[gx][gy] += "R"
				if strings.Contains(obstructed_map[gx][gy], "#") {
					d = direction(down)
					gy--
				}
			}
		case direction(down):
			gx++
			if gx >= len(obstructed_map) {
				return false
			}
			if strings.Contains(obstructed_map[gx][gy], "D") {
				return true
			} else {
				obstructed_map[gx][gy] += "D"
				if strings.Contains(obstructed_map[gx][gy], "#") {
					d = direction(left)
					gx--
				}
			}
		case direction(left):
			gy--
			if gy < 0 {
				return false
			}
			if strings.Contains(obstructed_map[gx][gy], "L") {
				return true
			} else {
				obstructed_map[gx][gy] += "L"
				if strings.Contains(obstructed_map[gx][gy], "#") {
					d = direction(up)
					gy++
				}
			}
		}
	}
}

func printGuardMap(guard_map [][]string) {
	for i := 0; i < len(guard_map); i++ {
		fmt.Println(guard_map[i])
	}
	fmt.Println("-------------------------------")
}
