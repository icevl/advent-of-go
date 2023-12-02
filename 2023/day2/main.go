package main

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	ID   int
	Sets []Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func (g *Game) getPower() int {
	var red, green, blue []int

	for _, set := range g.Sets {
		red = append(red, set.Red)
		green = append(green, set.Green)
		blue = append(blue, set.Blue)
	}

	return slices.Max(red) * slices.Max(green) * slices.Max(blue)
}

func main() {
	log.Println("Part 1:", part1())
	log.Println("Part 2:", part2())
}

func getGamesByLimit(games []Game, red int, green int, blue int) []Game {
	var possibleGames []Game

	for _, game := range games {
		var gameRed, gameGreen, gameBlue int
		var isPossible = true

		for _, set := range game.Sets {
			gameRed += set.Red
			gameGreen += set.Green
			gameBlue += set.Blue

			if set.Red > red || set.Green > green || set.Blue > blue {
				isPossible = false
				break
			}

		}

		if isPossible {
			possibleGames = append(possibleGames, game)
		}

	}

	return possibleGames
}

func loadGames() []Game {
	var games []Game
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		game, err := extractGame(line)
		if err != nil {
			continue
		}

		games = append(games, *game)
	}

	return games
}

func extractGame(data string) (*Game, error) {
	re := regexp.MustCompile(`Game\s(\d+):\s*(.*)`)
	matches := re.FindStringSubmatch(data)

	sets := []Set{}
	gameNumber := matches[1]
	gameDataStr := strings.Trim(matches[2], "")
	gameId, err := strconv.Atoi(gameNumber)

	if err != nil {
		return nil, err
	}

	setsStrings := strings.Split(gameDataStr, ";")

	for _, setString := range setsStrings {
		set := extractSet(setString)
		sets = append(sets, set)
	}

	game := Game{
		ID:   gameId,
		Sets: sets,
	}

	return &game, nil
}

func extractSet(data string) Set {
	var red, green, blue int

	params := strings.Split(data, ",")

	for _, param := range params {
		paramRe := regexp.MustCompile(`\b(?:red|green|blue)\b`)
		valueReg, _ := regexp.Compile("[^0-9]+")

		color := paramRe.FindAllString(param, -1)[0]
		valueStr := valueReg.ReplaceAllString(param, "")
		value, _ := strconv.Atoi(valueStr)

		switch color {
		case "red":
			red = value
		case "green":
			green = value
		case "blue":
			blue = value
		}

	}

	return Set{
		Red:   red,
		Green: green,
		Blue:  blue,
	}

}
