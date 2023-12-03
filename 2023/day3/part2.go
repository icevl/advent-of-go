package main

const MULTIPLY_SYMBOL = '*'

type Number struct {
	start int
	end   int
	y     int
	value int
}

func Part2(schematic Schematic) int {
	sum := 0
	allNumbers := make([]Number, 0)

	for y, line := range schematic {
		allNumbers = append(allNumbers, parseNumbersFromLine(string(line), y)...)
	}

	for y := 0; y < len(schematic); y++ {
		for x := 0; x < len(schematic[0]); x++ {
			if schematic[y][x] == MULTIPLY_SYMBOL {
				sum += getGear(x, y, allNumbers)
			}
		}
	}

	return sum
}

func parseNumbersFromLine(line string, y int) []Number {
	numbers := make([]Number, 0)

	for x := 0; x < len(line); x++ {
		start := x
		current := 0

		for ; x < len(line) && isDigit(line[x]); x++ {
			current = current*10 + int(line[x]-'0')
		}

		if current != 0 {
			x--
			numbers = append(numbers, Number{start, x, y, current})
		}
	}

	return numbers
}

func getGear(x int, y int, numbers []Number) int {
	validNumbers := getValidNumbers(numbers, x, y)

	if len(validNumbers) == 2 {
		return validNumbers[0].value * validNumbers[1].value
	}

	return 0
}

func getValidNumbers(numbers []Number, x int, y int) []Number {
	validNumbers := make([]Number, 0)

	for _, number := range numbers {
		if number.y == y {
			if number.start == x+1 || number.end == x-1 {
				validNumbers = append(validNumbers, number)
			}
			continue
		}

		if number.y == y+1 || number.y == y-1 {
			if number.start <= x+1 && number.end >= x-1 {
				validNumbers = append(validNumbers, number)
			}
		}
	}

	return validNumbers
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
