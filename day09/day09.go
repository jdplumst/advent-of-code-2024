package day09

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day09(part string) {
	switch part {
	case "1":
		res := part1("./day09/input.txt")
		fmt.Println("The result for day 9 part 1 is", res)
	case "2":
		res := part2("./day09/input.txt")
		fmt.Println("The result for day 9 part 2 is", res)
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

	blocks := make([]string, 0)

	file_block := true

	id := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			continue
		}
		blocks, file_block = createBlocks(blocks, scanner.Text(), id, file_block)

		if !file_block {
			id++
		}
	}

	blocks = moveBlocks(blocks)

	return getChecksum(blocks)
}

func part2(input string) int {
	return 1
}

func createBlocks(blocks []string, c string, id int, f bool) ([]string, bool) {
	x, err := strconv.Atoi(c)
	if err != nil {
		log.Fatal(err)
	}

	if f {
		for range x {
			blocks = append(blocks, strconv.Itoa(id))
		}
		return blocks, false
	} else {
		for range x {
			blocks = append(blocks, ".")
		}
		return blocks, true
	}
}

func moveBlocks(blocks []string) []string {
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == "." {
			blocks[i] = blocks[len(blocks)-1]
			blocks = blocks[:len(blocks)-1]
			blocks = removeTrailingSpace(blocks)
		}
	}

	return blocks
}

func removeTrailingSpace(blocks []string) []string {
	if blocks[len(blocks)-1] == "." {
		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i] != "." {
				break
			} else {
				blocks = blocks[:len(blocks)-1]
			}
		}
	}

	return blocks
}

func getChecksum(blocks []string) int {
	checksum := 0

	for i, e := range blocks {
		id, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal(err)
		}
		checksum += i * id
	}

	return checksum
}
