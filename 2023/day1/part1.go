package main

func part1() int {
	var sum int

	input := loadInput()

	for _, line := range input {
		sum += extractValue(line, false)
	}

	return sum
}
