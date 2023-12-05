package main

func part2() int {
	seeds, mappings := loadInput("input.txt")
	seedRanges := getSeedPairs(seeds)
	minLocation := solve2(seedRanges, mappings)
	return minLocation
}

func getSeedPairs(seeds []int) [][2]int {
	seedPairs := make([][2]int, 0, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		seedPairs = append(seedPairs, [2]int{seeds[i], seeds[i+1]})
	}
	return seedPairs
}

func solve2(seedRanges [][2]int, mappings Mapping) int {
	minLocation := int(1e9)

	for _, seedRange := range seedRanges {
		rangeStart, rangeLength := seedRange[0], seedRange[1]
		for i := 0; i < rangeLength; i++ {
			seed := rangeStart + i
			location := seed
			for _, mapping := range mappings {
				location = getEndPoint(location, mapping)
			}
			if location < minLocation {
				minLocation = location
			}
		}
	}

	return minLocation
}

func getEndPoint(value int, maps [][3]int) int {
	for _, values := range maps {
		destStart, srcStart, length := values[0], values[1], values[2]

		if value >= srcStart && value < srcStart+length {
			return destStart + (value - srcStart)
		}
	}
	return value
}
