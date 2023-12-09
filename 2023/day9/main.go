package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Println("Part 1:", part1("input.txt"))
	log.Println("Part 2:", part2("input.txt"))
}

func loadInput(file string) [][]int {

	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")

	var result [][]int

	for _, line := range lines {
		lineNumbers := getValues(line)
		result = append(result, lineNumbers)
	}

	return result
}

func getValues(line string) []int {
	var result []int
	for _, v := range strings.Split(line, " ") {
		valueInt, _ := strconv.Atoi(v)
		result = append(result, valueInt)
	}
	return result
}

func isAllZeroes(sequence []int) bool {
	for _, v := range sequence {
		if v != 0 {
			return false
		}
	}
	return true
}
