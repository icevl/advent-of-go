package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1Small(t *testing.T) {
	assert.Equal(t, part1("input_test.txt"), 6440)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, part1("input.txt"), 252295678)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, part2("input.txt"), 250577259)
}
