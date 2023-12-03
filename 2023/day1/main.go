package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var NUMBERS = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	log.Println("Part 1:", Part1())
	log.Println("Part 2:", Part2())
}

func loadInput() []string {
	input, _ := os.ReadFile("input.txt")
	return strings.Split(string(input), "\n")
}

func extractValue(input string, replaceWords bool) int {
	reg, _ := regexp.Compile("[^0-9]+")

	if replaceWords {
		input = wordsToNumbers(input)
	}

	stripNoneNumbers := reg.ReplaceAllString(input, "")

	first := string(stripNoneNumbers[0])
	last := string(stripNoneNumbers[len(stripNoneNumbers)-1])

	result, _ := strconv.Atoi(first + last)
	return result
}

func wordsToNumbers(input string) string {
	var buf string
	var resultStr []string

	for _, word := range strings.Split(input, "") {
		if unicode.IsNumber(rune(word[0])) {
			num := rune(word[0])
			resultStr = append(resultStr, string(num))

			buf = ""
			continue
		}

		buf += word

		for k := range NUMBERS {
			if strings.Contains(buf, k) {
				resultStr = append(resultStr, strconv.Itoa(NUMBERS[k]))
				buf = word
			}
		}

	}

	return strings.Join(resultStr, "")

}
