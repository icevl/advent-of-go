package main

import "log"

func part2() int {
	var sum int

	input, err := loadInput()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range input {
		sum += extractValue(line, true)
	}

	return sum
}
