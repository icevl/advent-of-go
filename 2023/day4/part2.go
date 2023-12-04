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

		for j := 0; j < deck[card.id]; j++ {
			for i := card.id + 1; i <= card.id+len(card.matched); i++ {
				deck[i] += 1
			}
		}

	}

	for _, value := range deck {
		sum += value
	}

	return sum
}
