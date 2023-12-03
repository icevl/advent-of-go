package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1(t *testing.T) {
	schematic := loadSchematic("input.txt")
	assert.Equal(t, Part1(schematic), 521601)
}

func TestPart2Small(t *testing.T) {
	schematic := loadSchematic("input_test.txt")
	assert.Equal(t, Part2(schematic), 467835)
}

func TestPart2(t *testing.T) {
	schematic := loadSchematic("input.txt")
	assert.Equal(t, Part2(schematic), 80694070)
}

func Benchmark_Part1(b *testing.B) {

	for n := 0; n < b.N; n++ {
		schematic := loadSchematic("input.txt")
		Part1(schematic)
	}
}

func Benchmark_Part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		schematic := loadSchematic("input.txt")
		Part2(schematic)
	}
}
