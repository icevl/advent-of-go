package main

func Part2() int {
	games := loadGames()

	var power int

	for _, game := range games {
		power += game.getPower()
	}

	return power
}
