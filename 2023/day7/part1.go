package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(file string) int {
	total := 0
	hands := loadInputPart1(file)

	sort.Slice(hands, func(i, j int) bool {
		return comparePower(hands[i], hands[j], "23456789TJQKA")
	})

	for rank := 1; rank <= len(hands); rank++ {
		total += rank * hands[rank-1].Bid
	}

	return total
}

func loadInputPart1(file string) []Hand {
	var hands []Hand

	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		hands = append(hands, parseInputPart1(line))
	}

	return hands
}

func parseInputPart1(line string) Hand {
	strs := strings.Split(line, " ")
	card := strs[0]
	bid, err := strconv.Atoi(strs[1])

	if err != nil {
		log.Panic(err)
	}

	m := map[rune]int{}
	for _, v := range card {
		m[v]++
	}

	first, firstAmount := getLargest(m, 0)
	_, secondAmount := getLargest(m, first)

	combo := HIGH_CARD
	switch firstAmount {

	case 5:
		combo = FIVE_OF_A_KIND

	case 4:
		combo = FOUR_OF_A_KIND

	case 3:
		combo = THREE_OF_A_KIND
		if secondAmount == 2 {
			combo = FULL_HOUSE
		}

	case 2:
		combo = ONE_PAIR
		if secondAmount == 2 {
			combo = TWO_PAIR
		}
	}

	return Hand{
		Card:  card,
		Bid:   bid,
		Combo: combo,
	}
}
