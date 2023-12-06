package main

func part1(file string) int {
	result := 1

	races := loadRaces(file, false)

	for _, race := range races {
		variants := race.getWinVariants()
		result *= variants
	}

	return result
}
