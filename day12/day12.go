package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day12(part string) {
	switch part {
	case "1":
		res := part1("./day12/input.txt")
		fmt.Println("The result for day 12 part 1 is", res)
	case "2":
		res := part2("./day12/input.txt")
		fmt.Println("The result for day 12 part 2 is", res)
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

	garden := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), "")
		garden = append(garden, x)
	}

	tracker := make([][]bool, len(garden))
	for i := range tracker {
		tracker[i] = make([]bool, len(garden[i]))
	}

	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			if !tracker[i][j] {
				a, p := trackRegion(garden, tracker, i, j)
				res += a * p
			}
		}
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

	garden := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), "")
		garden = append(garden, x)
	}

	tracker := make([][]bool, len(garden))
	for i := range tracker {
		tracker[i] = make([]bool, len(garden[i]))
	}

	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			if !tracker[i][j] {
				cornerTracker := make([][]int, len(garden)+1)
				for i := range cornerTracker {
					cornerTracker[i] = make([]int, len(garden)+1)
				}
				a, cornerTracker := trackRegionUpdated(garden, tracker, cornerTracker, i, j)
				c := countConers(cornerTracker)
				res += a * c
			}
		}
	}

	return res
}

func trackRegion(garden [][]string, tracker [][]bool, i, j int) (int, int) {
	tracker[i][j] = true
	area := 1
	perimeter := 4

	// Check north
	if i > 0 && garden[i-1][j] == garden[i][j] {
		perimeter--
		if !tracker[i-1][j] {
			a, p := trackRegion(garden, tracker, i-1, j)
			area += a
			perimeter += p
		}
	}

	// Check south
	if i < len(garden)-1 && garden[i+1][j] == garden[i][j] {
		perimeter--
		if !tracker[i+1][j] {
			a, p := trackRegion(garden, tracker, i+1, j)
			area += a
			perimeter += p
		}
	}

	// Check west
	if j > 0 && garden[i][j-1] == garden[i][j] {
		perimeter--
		if !tracker[i][j-1] {
			a, p := trackRegion(garden, tracker, i, j-1)
			area += a
			perimeter += p
		}
	}

	// Check east
	if j < len(garden[i])-1 && garden[i][j+1] == garden[i][j] {
		perimeter--
		if !tracker[i][j+1] {
			a, p := trackRegion(garden, tracker, i, j+1)
			area += a
			perimeter += p
		}
	}

	return area, perimeter
}

func trackRegionUpdated(garden [][]string, tracker [][]bool, cornerTracker [][]int, i, j int) (int, [][]int) {
	tracker[i][j] = true
	area := 1

	// Check north
	if i > 0 && garden[i-1][j] == garden[i][j] {
		if !tracker[i-1][j] {
			a, _ := trackRegionUpdated(garden, tracker, cornerTracker, i-1, j)
			area += a
		}
	}

	// Check south
	if i < len(garden)-1 && garden[i+1][j] == garden[i][j] {
		if !tracker[i+1][j] {
			a, _ := trackRegionUpdated(garden, tracker, cornerTracker, i+1, j)
			area += a
		}
	}

	// Check west
	if j > 0 && garden[i][j-1] == garden[i][j] {
		if !tracker[i][j-1] {
			a, _ := trackRegionUpdated(garden, tracker, cornerTracker, i, j-1)
			area += a
		}
	}

	// Check east
	if j < len(garden[i])-1 && garden[i][j+1] == garden[i][j] {
		if !tracker[i][j+1] {
			a, _ := trackRegionUpdated(garden, tracker, cornerTracker, i, j+1)
			area += a
		}
	}

	// Check top-left corner
	if i == 0 && j == 0 {
		cornerTracker[i][j] = 1
	} else if i > 0 && garden[i-1][j] != garden[i][j] &&
		j > 0 && garden[i][j-1] != garden[i][j] &&
		garden[i-1][j-1] == garden[i][j] && cornerTracker[i][j] != 0 {
		cornerTracker[i][j] = 2
	} else if ((i > 0 && garden[i-1][j] != garden[i][j]) || (i == 0)) &&
		((j > 0 && garden[i][j-1] != garden[i][j]) || (j == 0)) {
		cornerTracker[i][j] = 1
	} else if i > 0 && garden[i-1][j] == garden[i][j] &&
		j > 0 && garden[i][j-1] == garden[i][j] &&
		garden[i-1][j-1] != garden[i][j] {
		cornerTracker[i][j] = 1
	}

	// Check top-right corner
	if i == 0 && j == len(garden[i])-1 {
		cornerTracker[i][j+1] = 1
	} else if i > 0 && garden[i-1][j] != garden[i][j] &&
		j < len(garden[i])-1 && garden[i][j+1] != garden[i][j] &&
		garden[i-1][j+1] == garden[i][j] && cornerTracker[i][j+1] != 0 {
		cornerTracker[i][j+1] = 2
	} else if ((i > 0 && garden[i-1][j] != garden[i][j]) || (i == 0)) &&
		((j < len(garden[i])-1 && garden[i][j+1] != garden[i][j]) || (j == len(garden[i])-1)) {
		cornerTracker[i][j+1] = 1
	} else if i > 0 && garden[i-1][j] == garden[i][j] &&
		j < len(garden[i])-1 && garden[i][j+1] == garden[i][j] &&
		garden[i-1][j+1] != garden[i][j] {
		cornerTracker[i][j+1] = 1
	}

	// Check bot-left corner
	if i == len(garden)-1 && j == 0 {
		cornerTracker[i+1][j] = 1
	} else if i < len(garden)-1 && garden[i+1][j] != garden[i][j] &&
		j > 0 && garden[i][j-1] != garden[i][j] &&
		garden[i+1][j-1] == garden[i][j] && cornerTracker[i+1][j] != 0 {
		cornerTracker[i+1][j] = 2
	} else if ((i < len(garden)-1 && garden[i+1][j] != garden[i][j]) || (i == len(garden)-1)) &&
		((j > 0 && garden[i][j-1] != garden[i][j]) || (j == 0)) {
		cornerTracker[i+1][j] = 1
	} else if i < len(garden)-1 && garden[i+1][j] == garden[i][j] &&
		j > 0 && garden[i][j-1] == garden[i][j] &&
		garden[i+1][j-1] != garden[i][j] {
		cornerTracker[i+1][j] = 1
	}

	// Check bot-right corner
	if i == len(garden)-1 && j == len(garden[i])-1 {
		cornerTracker[i+1][j+1] = 1
	} else if i < len(garden)-1 && garden[i+1][j] != garden[i][j] &&
		j < len(garden[i])-1 && garden[i][j+1] != garden[i][j] &&
		garden[i+1][j+1] == garden[i][j] && cornerTracker[i+1][j+1] != 0 {
		cornerTracker[i+1][j+1] = 2
	} else if ((i < len(garden)-1 && garden[i+1][j] != garden[i][j]) || (i == len(garden)-1)) &&
		((j < len(garden)-1 && garden[i][j+1] != garden[i][j]) || (j == len(garden)-1)) {
		cornerTracker[i+1][j+1] = 1
	} else if i < len(garden)-1 && garden[i+1][j] == garden[i][j] &&
		j < len(garden[i])-1 && garden[i][j+1] == garden[i][j] &&
		garden[i+1][j+1] != garden[i][j] {
		cornerTracker[i+1][j+1] = 1
	}

	return area, cornerTracker
}

func countConers(cornerTracker [][]int) int {
	corners := 0

	for i := range cornerTracker {
		for j := range cornerTracker[i] {
			corners += cornerTracker[i][j]
		}
	}

	return corners
}
