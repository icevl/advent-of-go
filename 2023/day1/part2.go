package main

func part2() int {
	var sum int

	input := loadInput()

	for _, line := range input {
		sum += extractValue(line, true)
	}

	return sum
}
