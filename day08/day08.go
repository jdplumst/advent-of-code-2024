package day08

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day08(part string) {
	switch part {
	case "1":
		res := part1("./day08/input.txt")
		fmt.Println("The result for day 8 part 1 is", res)
	case "2":
		res := part2("./day08/input.txt")
		fmt.Println("The result for day 8 part 2 is", res)
	default:
		log.Fatal("Must run either part 1 or part 2")
	}
}

type Coordinate struct {
	X int
	Y int
}

func part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	antennas := map[string][]Coordinate{}

	width := 0
	height := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), "")
		width = len(c)
		for i := range c {
			if c[i] != "." {
				index := c[i]
				antennas[index] = append(antennas[index], Coordinate{i, height})
			}
		}
		height++
	}

	antinodes := make([]Coordinate, 0)
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				new_X := abs(v[i].X, v[j].X)
				new_Y := abs(v[i].Y, v[j].Y)

				// i top-left, j bot-right
				if v[i].X <= v[j].X && v[i].Y <= v[j].Y {
					first_X := v[i].X - new_X
					first_Y := v[i].Y - new_Y
					if first_X >= 0 && first_X < width &&
						first_Y >= 0 && first_Y < height &&
						!checkExists(antinodes, first_X, first_Y) {
						antinodes = append(antinodes, Coordinate{first_X, first_Y})
					}

					second_X := v[j].X + new_X
					second_Y := v[j].Y + new_Y
					if second_X >= 0 && second_X < width &&
						second_Y >= 0 && second_Y < height &&
						!checkExists(antinodes, second_X, second_Y) {
						antinodes = append(antinodes, Coordinate{second_X, second_Y})
					}
					// i top-right, j bot-left
				} else if v[i].X >= v[j].X && v[i].Y <= v[j].Y {
					first_X := v[i].X + new_X
					first_Y := v[i].Y - new_Y
					if first_X >= 0 && first_X < width &&
						first_Y >= 0 && first_Y < height &&
						!checkExists(antinodes, first_X, first_Y) {
						antinodes = append(antinodes, Coordinate{first_X, first_Y})
					}

					second_X := v[j].X - new_X
					second_Y := v[j].Y + new_Y
					if second_X >= 0 && second_X < width &&
						second_Y >= 0 && second_Y < height &&
						!checkExists(antinodes, second_X, second_Y) {
						antinodes = append(antinodes, Coordinate{second_X, second_Y})
					}
					// i bot-left, j top-right
				} else if v[i].X <= v[j].X && v[i].Y >= v[j].Y {
					first_X := v[i].X - new_X
					first_Y := v[i].Y + new_Y
					if first_X >= 0 && first_X < width &&
						first_Y >= 0 && first_Y < height &&
						!checkExists(antinodes, first_X, first_Y) {
						antinodes = append(antinodes, Coordinate{first_X, first_Y})
					}

					second_X := v[j].X + new_X
					second_Y := v[j].Y - new_Y
					if second_X >= 0 && second_X < width &&
						second_Y >= 0 && second_Y < height &&
						!checkExists(antinodes, second_X, second_Y) {
						antinodes = append(antinodes, Coordinate{second_X, second_Y})
					}
					// i bot-right, j top-left
				} else if v[i].X >= v[j].X && v[i].Y >= v[j].Y {
					first_X := v[i].X + new_X
					first_Y := v[i].Y + new_Y
					if first_X >= 0 && first_X < width &&
						first_Y >= 0 && first_Y < height &&
						!checkExists(antinodes, first_X, first_Y) {
						antinodes = append(antinodes, Coordinate{first_X, first_Y})
					}

					second_X := v[j].X - new_X
					second_Y := v[j].Y - new_Y
					if second_X >= 0 && second_X < width &&
						second_Y >= 0 && second_Y < height &&
						!checkExists(antinodes, second_X, second_Y) {
						antinodes = append(antinodes, Coordinate{second_X, second_Y})
					}
				}
			}
		}
	}

	return len(antinodes)
}

func part2(input string) int {
	return 1
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

func checkExists(a []Coordinate, x, y int) bool {
	for i := range a {
		if a[i].X == x && a[i].Y == y {
			return true
		}
	}
	return false
}

// 307 too high
// 300 too low
