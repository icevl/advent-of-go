package main

import (
	"fmt"
)

func part2() (int, error) {
	pipeMap, err := loadInput()
	if err != nil {
		return 0, fmt.Errorf("could not read input: %w", err)
	}

	loop, err := findLoop(pipeMap)
	if err != nil {
		return 0, fmt.Errorf("could not find loop: %w", err)
	}

	area := areaWithinLoop(pipeMap, loop)
	return area, nil
}

/*
- Zooms in on the map (represented as zoomedInMap) by a factor of 2
to better define the outer boundary of the loop.
- Uses a stack to find the area within the loop.
- The area within the loop is the area of the map minus the area outside the loop.
- The area outside the loop is the area of the map minus the area inside the loop.
*/
func areaWithinLoop(pipeMap [][]byte, loop []Position) int {
	zoomedInMap := make([][]byte, len(pipeMap)*2+1)

	for i := range zoomedInMap {
		zoomedInMap[i] = make([]byte, len(pipeMap[0])*2+1)
	}

	for row := range zoomedInMap {
		for col := range zoomedInMap[row] {
			zoomedInMap[row][col] = ground
		}
	}

	for i := range loop {
		pos, nextPos := loop[i], loop[(i+1)%len(loop)]

		zoomedInMap[pos.row*2+1][pos.col*2+1] = pipeMap[pos.row][pos.col]

		rowDelta, colDelta := nextPos.row-pos.row, nextPos.col-pos.col

		switch {
		case rowDelta == -1 && colDelta == 0:
			zoomedInMap[pos.row*2][pos.col*2+1] = pipeVertical
		case rowDelta == 1 && colDelta == 0:
			zoomedInMap[pos.row*2+2][pos.col*2+1] = pipeVertical
		case rowDelta == 0 && colDelta == -1:
			zoomedInMap[pos.row*2+1][pos.col*2] = pipeHorizontal
		case rowDelta == 0 && colDelta == 1:
			zoomedInMap[pos.row*2+1][pos.col*2+2] = pipeHorizontal
		default:
			panic("diagonal pipe")
		}
	}

	outsideArea := 0

	seen := make([][]bool, len(zoomedInMap))
	for i := range seen {
		seen[i] = make([]bool, len(zoomedInMap[i]))
	}

	var stack []Position

	processPosition := func(pos Position) {
		if pos.row < 0 || pos.row >= len(zoomedInMap) || pos.col < 0 || pos.col >= len(zoomedInMap[pos.row]) {
			return
		}
		if zoomedInMap[pos.row][pos.col] != ground {
			return
		}
		if seen[pos.row][pos.col] {
			return
		}

		stack = append(stack, pos)
		seen[pos.row][pos.col] = true

		if pos.row%2 == 1 && pos.col%2 == 1 {
			outsideArea++
		}
	}

	pos := Position{0, 0}
	stack = append(stack, pos)
	seen[pos.row][pos.col] = true

	for len(stack) > 0 {
		pos, stack = stack[len(stack)-1], stack[:len(stack)-1]

		processPosition(Position{pos.row - 1, pos.col})
		processPosition(Position{pos.row + 1, pos.col})
		processPosition(Position{pos.row, pos.col - 1})
		processPosition(Position{pos.row, pos.col + 1})
	}

	mapArea := len(pipeMap) * len(pipeMap[0])
	insideArea := mapArea - outsideArea - len(loop)

	return insideArea
}
