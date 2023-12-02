package main

func part1() int {
	games := loadGames()

	var sum int
	limitedGames := getGamesByLimit(games, 12, 13, 14)

	for _, game := range limitedGames {
		sum += game.ID
	}

	return sum
}
