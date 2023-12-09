package main

func part1(file string) int {
	histories := loadInput(file)
	sum := 0

	for _, sequence := range histories {
		sum += nextValue(sequence)
	}

	return sum
}

func nextValue(sequence []int) int {
	if isAllZeroes(sequence) {
		return 0
	}

	d := diffs(sequence)
	nextDiff := nextValue(d)
	return sequence[len(sequence)-1] + nextDiff
}

func diffs(sequence []int) []int {
	diffs := make([]int, len(sequence)-1)

	for i := 0; i < len(sequence)-1; i++ {
		diffs[i] = sequence[i+1] - sequence[i]
	}

	return diffs
}
