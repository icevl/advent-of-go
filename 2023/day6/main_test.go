package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1Small(t *testing.T) {
	assert.Equal(t, part1("input_test.txt"), 288)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, part1("input.txt"), 781200)
}

func TestPart2Small(t *testing.T) {
	assert.Equal(t, part2("input_test.txt"), 71503)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, part2("input.txt"), 49240091)
}

func Benchmark_Part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part1("input.txt")
	}
}

func Benchmark_Part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part2("input.txt")
	}
}
