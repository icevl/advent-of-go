package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1(t *testing.T) {
	schematic := loadSchematic("input.txt")
	assert.Equal(t, part1(schematic), 521601)
}

func TestPart2Small(t *testing.T) {
	schematic := loadSchematic("input_test.txt")
	assert.Equal(t, part2(schematic), 467835)
}

func TestPart2(t *testing.T) {
	schematic := loadSchematic("input.txt")
	assert.Equal(t, part2(schematic), 80694070)
}
