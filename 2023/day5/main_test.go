package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, part1(), 825516882)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, part2(), 136096660)
}
