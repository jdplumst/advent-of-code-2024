package day07

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day07(part string) {
	switch part {
	case "1":
		res := part1("./day07/input.txt")
		fmt.Println("The result for day 7 part 1 is", res)
	case "2":
		res := part2("./day07/input.txt")
		fmt.Println("The result for day 7 part 2 is", res)
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
		x := strings.Split(scanner.Text(), " ")
		x[0] = strings.TrimSuffix(x[0], ":")
		y := make([]int, len(x))
		for i, e := range x {
			v, err := strconv.Atoi(e)
			if err != nil {
				log.Fatal(err)
			}
			y[i] = v
		}

		if checkValid(y) {
			res += y[0]
		}
	}

	return res
}

func part2(input string) int {
	return 2
}

func checkValid(a []int) bool {
	operators := []string{"+", "*"}

	tree := make([]int, 0)

	tree = append(tree, a[1])

	total_nodes := int(math.Pow(float64(2), float64(len(a)-1))) - 1

	depth := 1
	parentIndex := 0
	for i := 1; i <= (total_nodes / 2); i++ {
		for _, j := range operators {
			if j == "+" {
				tree = append(tree, tree[parentIndex]+a[depth+1])
			} else {
				tree = append(tree, tree[parentIndex]*a[depth+1])
			}
		}
		parentIndex++
		depth = getDepth(tree, len(tree))
	}

	return slices.Contains(tree, a[0])
}

func getDepth(tree []int, index int) int {
	if index == 0 {
		return 0
	}
	return 1 + getDepth(tree, (index-1)/2)

}
