package main

func Part1() int {
	sum := 0

	cards := loadCards("input.txt")

	for i := range cards {
		sum += cards[i].GetPoints()
	}

	return sum
}
