package main

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id         int
	numbers    []int
	winNumbers []int
	matched    []int
}

func (c *Card) GetPoints() int {
	matchedCount := len(c.matched)

	if matchedCount == 0 {
		return 0
	}

	return 1 << (matchedCount - 1)
}

func main() {
	log.Println("Part 1:", Part1())
	log.Println("Part 2:", Part2())
}

func loadCards(file string) []Card {
	var cards []Card

	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")

	for row := range lines {
		if card, ok := extractCard(lines[row]); ok {
			cards = append(cards, *card)
		}
	}

	return cards
}

func extractCard(line string) (*Card, bool) {
	re := regexp.MustCompile(`Card\s+(\d+):\s+(.*)\s+\|\s+(.*)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 3 {
		return nil, false
	}

	id, _ := strconv.Atoi(matches[1])
	winNumbers := extractNumbers(matches[2])
	numbers := extractNumbers(matches[3])

	var matched []int

	for _, number := range winNumbers {
		if slices.Contains(numbers, number) {
			matched = append(matched, number)
		}
	}

	card := Card{
		id, numbers, winNumbers, matched,
	}

	return &card, true

}

func extractNumbers(line string) []int {
	var numbers []int

	re := regexp.MustCompile(`\d{1,3}`)
	matches := re.FindAllStringSubmatch(line, -1)

	for i := range matches {
		number, _ := strconv.Atoi(matches[i][0])
		numbers = append(numbers, number)
	}

	return numbers
}
