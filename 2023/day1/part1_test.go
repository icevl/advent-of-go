package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1Extraction(t *testing.T) {
	n1 := extractValue("1abc2", false)
	n2 := extractValue("pqr3stu8vwx", false)
	n3 := extractValue("a1b2c3d4e5f", false)
	n4 := extractValue("treb7uchet", false)

	assert.Equal(t, n1, 12)
	assert.Equal(t, n2, 38)
	assert.Equal(t, n3, 15)
	assert.Equal(t, n4, 77)

	assert.Equal(t, n1+n2+n3+n4, 142)
}

func TestResultPart1(t *testing.T) {
	assert.Equal(t, part1(), 53386)
}
