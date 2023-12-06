package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type race struct {
	time           int
	recordDistance int
}

func main() {
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Part 2:", part2("input.txt"))
}

func (r *race) getWinVariants() int {
	variants := 0

	for holdTime := 0; holdTime <= r.time; holdTime++ {

		travelTime := r.time - holdTime
		distance := holdTime * travelTime

		if distance > r.recordDistance {
			variants += 1
		}
	}
	return variants
}

func loadRaces(file string, concatValues bool) []race {
	var races []race

	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")

	var times []int
	var distances []int

	for _, line := range lines {
		if strings.HasPrefix(line, "Time:") {
			if concatValues {
				times = append(times, getSingleValue(line))
			} else {
				times = getValues(line)
			}
		}

		if strings.HasPrefix(line, "Distance:") {
			if concatValues {
				distances = append(distances, getSingleValue(line))
			} else {
				distances = getValues(line)
			}
		}
	}

	if len(times) != len(distances) {
		fmt.Println("Amount of times and distances are not equal")
		os.Exit(1)
	}

	for i := range times {
		time := times[i]
		recordDistance := distances[i]
		races = append(races, race{time, recordDistance})
	}

	return races
}

func getValues(line string) []int {
	var numbers []int

	stripRegex := regexp.MustCompile(`[^0-9 ]+`)
	seedsExp := regexp.MustCompile(`\d+`)

	stripedLine := stripRegex.ReplaceAllString(line, "")
	list := seedsExp.FindAllString(stripedLine, -1)

	for _, values := range list {
		valueInt, _ := strconv.Atoi(values)
		numbers = append(numbers, valueInt)
	}

	return numbers
}

func getSingleValue(line string) int {
	number := 0
	value := strings.ReplaceAll(line[strings.Index(line, ":")+1:], " ", "")
	number, _ = strconv.Atoi(value)
	return number
}
