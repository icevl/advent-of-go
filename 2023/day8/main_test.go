package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1Small(t *testing.T) {
	assert.Equal(t, part1("input_test.txt"), 2)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, part1("input.txt"), 16897)
}

func TestPart2Small(t *testing.T) {
	assert.Equal(t, part2("input_test2.txt"), 6)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, part2("input.txt"), 16563603485021)
}
