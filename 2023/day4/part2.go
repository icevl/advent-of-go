package main

func Part2() int {
	cards := loadCards("input.txt")

	sum := 0
	deck := map[int]int{}

	for i := range cards {
		card := cards[i]
		deck[card.id] += 1

		if len(cards[i].matched) == 0 {
			continue
		}

		for d := 0; d < deck[card.id]; d++ {
			for j := card.id + 1; j <= card.id+len(card.matched); j++ {
				deck[j] += 1
			}
		}

	}

	for _, value := range deck {
		sum += value
	}

	return sum
}
