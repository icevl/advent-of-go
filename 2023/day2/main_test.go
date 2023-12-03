package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestGameParser(t *testing.T) {
	game, _ := extractGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

	assert.Equal(t, len(game.Sets), 3)

	assert.Equal(t, game.Sets[0].Red, 4)
	assert.Equal(t, game.Sets[0].Blue, 3)
	assert.Equal(t, game.Sets[0].Green, 0)

	assert.Equal(t, game.Sets[1].Red, 1)
	assert.Equal(t, game.Sets[1].Green, 2)
	assert.Equal(t, game.Sets[1].Blue, 6)

	assert.Equal(t, game.Sets[2].Red, 0)
	assert.Equal(t, game.Sets[2].Green, 2)
	assert.Equal(t, game.Sets[2].Blue, 0)
}

func TestPowerCalculation(t *testing.T) {
	game1, _ := extractGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	game2, _ := extractGame("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")
	game3, _ := extractGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	game4, _ := extractGame("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")
	game5, _ := extractGame("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")

	assert.Equal(t, game1.getPower(), 48)
	assert.Equal(t, game2.getPower(), 12)
	assert.Equal(t, game3.getPower(), 1560)
	assert.Equal(t, game4.getPower(), 630)
	assert.Equal(t, game5.getPower(), 36)

}

func TestPart1(t *testing.T) {
	assert.Equal(t, Part1(), 2913)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, Part2(), 55593)
}
