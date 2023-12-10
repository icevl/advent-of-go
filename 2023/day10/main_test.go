package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1(t *testing.T) {
	result, _ := part1()
	assert.Equal(t, result, 6856)
}

func TestPart2(t *testing.T) {
	result, _ := part2()
	assert.Equal(t, result, 501)
}
