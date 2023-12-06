package main

func part2(file string) int {
	result := 1

	races := loadRaces(file, true)

	for _, race := range races {
		variants := race.getWinVariants()
		result *= variants
	}

	return result
}
