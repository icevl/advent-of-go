package main

import (
	"log"
)

const (
	start  = "AAA"
	target = "ZZZ"
)

func part1(file string) int {
	direction, desertMap := loadInput(file)

	current := start
	count := 0

	for {

		step := direction[count%len(direction)]
		next := getNextPoint(desertMap, current, step)

		if next == target {
			break
		}

		current = next
		count += 1
	}

	return count + 1
}

func getNextPoint(desertMap DesertMap, point string, direction string) string {
	switch direction {
	case "L":
		return desertMap[point][0]
	case "R":
		return desertMap[point][1]
	default:
		log.Panicln("Invalid direction")
		return ""
	}
}
