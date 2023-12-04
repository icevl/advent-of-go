package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1Short(t *testing.T) {
	card1, _ := extractCard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	card2, _ := extractCard("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19")
	card3, _ := extractCard("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1")
	card4, _ := extractCard("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83")
	card5, _ := extractCard("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36")
	card6, _ := extractCard("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11")

	assert.Equal(t, card1.GetPoints(), 8)
	assert.Equal(t, card2.GetPoints(), 2)
	assert.Equal(t, card3.GetPoints(), 2)
	assert.Equal(t, card4.GetPoints(), 1)
	assert.Equal(t, card5.GetPoints(), 0)
	assert.Equal(t, card6.GetPoints(), 0)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, Part1(), 24733)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, Part2(), 5422730)
}

func Benchmark_Part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1()
	}
}

func Benchmark_Part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part2()
	}
}
