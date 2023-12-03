package main

import (
	"log"
	"os"
	"strings"
)

type Schematic [][]rune

func main() {
	schematic := loadSchematic("input.txt")

	log.Println("Part 1:", part1(schematic))
	log.Println("Part 2:", part2(schematic))

}

func loadSchematic(file string) Schematic {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	grid := make([][]rune, len(lines))

	for y := range lines {
		grid[y] = make([]rune, len(lines[y]))
		for x, symbol := range lines[y] {
			grid[y][x] = symbol
		}
	}

	return grid
}
