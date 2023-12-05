package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mapping = [][][3]int

func main() {
	log.Println("Part 1:", part1())
	log.Println("Part 2:", part2())
}

func loadInput(file string) ([]int, Mapping) {
	var seeds []int
	var mapping Mapping

	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")

	mapBlock := false
	currentMap := [][3]int{}

	for i, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			seeds = getNumbers(line)
			continue
		}

		if line == "" && len(currentMap) > 0 {
			mapBlock = false
			mapping = append(mapping, currentMap)
			continue
		}

		if strings.Contains(line, "map:") {
			mapBlock = true
			currentMap = [][3]int{}
			continue
		}

		if mapBlock {
			numbers := getNumbers(line)
			destStart, srcStart, length := numbers[0], numbers[1], numbers[2]
			values := [3]int{destStart, srcStart, length}
			currentMap = append(currentMap, values)
		}

		if i == len(lines)-1 {
			mapping = append(mapping, currentMap)
		}
	}

	return seeds, mapping
}

func getNumbers(line string) []int {
	var numbers []int

	seedsExp := regexp.MustCompile(`\d+`)
	list := seedsExp.FindAllString(line, -1)

	for _, seed := range list {
		seedInt, _ := strconv.Atoi(seed)
		numbers = append(numbers, seedInt)
	}

	return numbers
}
