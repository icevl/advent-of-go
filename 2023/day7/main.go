package main

import (
	"log"
)

type Combo int

const (
	FIVE_OF_A_KIND  Combo = 7
	FOUR_OF_A_KIND  Combo = 6
	FULL_HOUSE      Combo = 5
	THREE_OF_A_KIND Combo = 4
	TWO_PAIR        Combo = 3
	ONE_PAIR        Combo = 2
	HIGH_CARD       Combo = 1
)

type Hand struct {
	Card  string
	Bid   int
	Combo Combo
}

func main() {
	log.Println("Part 1:", part1("input.txt"))
	log.Println("Part 2:", part2("input.txt"))
}

func getLargest(m map[rune]int, exclude rune) (rune, int) {
	maxAmount := 0
	var largest rune

	for k, v := range m {
		if k == exclude {
			continue
		}

		if v > maxAmount {
			maxAmount = v
			largest = k
		}
	}

	return largest, maxAmount
}

func comparePower(hand1 Hand, hand2 Hand, rules string) bool {
	findPos := func(c rune) int {
		for i, v := range rules {
			if c == v {
				return i
			}
		}

		return -1
	}

	if hand1.Combo == hand2.Combo {
		for i, v1 := range hand1.Card {
			v2 := rune(hand2.Card[i])
			if v1 == v2 {
				continue
			}

			return findPos(v1) < findPos(v2)
		}

		return true
	}

	return hand1.Combo < hand2.Combo

}
