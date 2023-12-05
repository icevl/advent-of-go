package main

func part1() int {
	seeds, mappings := loadInput("input.txt")
	result := solve(seeds, mappings)

	lowestLocationNumber := result[0]

	for _, r := range result {
		if r < lowestLocationNumber {
			lowestLocationNumber = r
		}
	}

	return lowestLocationNumber
}

func convert(mapping [][3]int, value int) int {
	for _, m := range mapping {
		destStart, srcStart, length := m[0], m[1], m[2]
		if srcStart <= value && value < srcStart+length {
			return destStart + (value - srcStart)
		}
	}
	return value
}

func solve(seedCodes []int, mappings Mapping) []int {
	results := []int{}

	for _, seedCode := range seedCodes {

		currentCode := seedCode

		for _, mapping := range mappings {
			currentCode = convert(mapping, currentCode)
		}

		results = append(results, currentCode)
	}

	return results
}
