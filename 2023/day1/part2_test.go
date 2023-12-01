package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart2Extraction(t *testing.T) {
	n1 := extractValue("two1nine", true)
	n2 := extractValue("eightwothree", true)
	n3 := extractValue("abcone2threexyz", true)
	n4 := extractValue("xtwone3four", true)
	n5 := extractValue("4nineeightseven2", true)
	n6 := extractValue("zoneight234", true)
	n7 := extractValue("7pqrstsixteen", true)

	assert.Equal(t, extractValue("six7sixqrdfive3twonehsk", true), 61)

	assert.Equal(t, n1, 29)
	assert.Equal(t, n2, 83)
	assert.Equal(t, n3, 13)
	assert.Equal(t, n4, 24)
	assert.Equal(t, n5, 42)
	assert.Equal(t, n6, 14)
	assert.Equal(t, n7, 76)

	assert.Equal(t, n1+n2+n3+n4+n5+n6+n7, 281)
}

func TestResultPart2(t *testing.T) {
	assert.Equal(t, part2(), 53312)
}
