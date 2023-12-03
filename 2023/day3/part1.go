package main

import (
	"strconv"
	"unicode"
)

const EMPTY_SYMBOL = '.'

func Part1(schematic Schematic) int {
	sum := 0

	for y := range schematic {
		var numStartPos = -1
		var numberCollected bool
		var collectedNumberChars string

		for x := range schematic[y] {
			char := schematic[y][x]

			if unicode.IsNumber(char) {
				if numStartPos == -1 {
					numStartPos = x
					collectedNumberChars = ""
					numberCollected = false
				}

				collectedNumberChars += string(char)

				if x == len(schematic[y])-1 {
					numberCollected = true
				}

			}

			if !unicode.IsNumber(char) && numStartPos != -1 {
				numberCollected = true
			}

			if numberCollected {
				num, _ := strconv.Atoi(collectedNumberChars)

				if hasNeighbors(schematic, numStartPos, y, num) {
					sum += num
				}

				collectedNumberChars = ""
				numStartPos = -1
			}

		}

	}

	return sum
}

func hasNeighbors(schematic Schematic, x int, y int, num int) bool {
	numLen := len(strconv.Itoa(num))

	if x-1 >= 0 && schematic[y][x-1] != EMPTY_SYMBOL {
		return true
	}

	if x+numLen < len(schematic[y]) && schematic[y][x+numLen] != EMPTY_SYMBOL {
		return true
	}

	if y-1 >= 0 {
		for i := x - 1; i < x+numLen+1; i++ {
			if i >= 0 && i < len(schematic[y-1]) && schematic[y-1][i] != EMPTY_SYMBOL {
				return true
			}
		}
	}

	if y+1 < len(schematic) {
		for i := x - 1; i < x+numLen+1; i++ {
			if i >= 0 && i < len(schematic[y+1]) && schematic[y+1][i] != EMPTY_SYMBOL {
				return true
			}
		}
	}

	return false
}
