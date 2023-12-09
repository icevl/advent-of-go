package main

func part2(file string) int {
	histories := loadInput(file)
	sum := 0

	for _, sequence := range histories {
		sum += previousValue(sequence)
	}

	return sum
}

func previousValue(sequence []int) int {
	if isAllZeroes(sequence) {
		return 0
	}

	d := diffs(sequence)
	previousDiff := previousValue(d)
	return sequence[0] - previousDiff
}
