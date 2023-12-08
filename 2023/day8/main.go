package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type Direction []string
type DesertMap map[string][]string

func main() {
	log.Println("Part 1:", part1("input.txt"))
	log.Println("Part 2:", part2("input.txt"))

}

func loadInput(file string) (Direction, DesertMap) {
	direction := make(Direction, 0)
	desertMap := make(DesertMap)

	re := regexp.MustCompile(`^(\w+)\s*=\s*\((\w+),\s*(\w+)\)`)
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")

	for index, line := range lines {
		if index == 0 {
			direction = strings.Split(line, "")
			continue
		}

		if strings.Contains(line, "=") {
			matches := re.FindStringSubmatch(line)
			if len(matches) < 4 {
				log.Panicln("Invalid input")
			}

			point, left, right := matches[1], matches[2], matches[3]

			desertMap[point] = []string{left, right}
		}
	}

	return direction, desertMap
}
