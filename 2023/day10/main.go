package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	pipeVertical      = '|'
	pipeHorizontal    = '-'
	pipeBendNorthEast = 'L'
	pipeBendNorthWest = 'J'
	pipeBendSouthWest = '7'
	pipeBendSouthEast = 'F'
	ground            = '.'
	startingPosition  = 'S'
)

type Position struct {
	row, col int
}

func main() {

	part1, err1 := part1()

	if err1 != nil {
		log.Fatal(err1)
	}

	p2, err2 := part2()

	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println("Part1:", part1)
	log.Println("Part2:", p2)
}

func loadInput() ([][]byte, error) {

	input, err := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	if err != nil {
		return nil, fmt.Errorf("error in file reading")
	}

	pipeMap := make([][]byte, len(lines))

	for i, line := range lines {
		pipeMap[i] = []byte(line)
	}

	if len(pipeMap) == 0 {
		return nil, fmt.Errorf("no input data")
	}

	lineLength := len(pipeMap[0])

	for _, line := range pipeMap {
		if len(line) != lineLength {
			return nil, fmt.Errorf("input is not rectangular")
		}
	}

	for row := range pipeMap {
		for col := range pipeMap[row] {
			switch pipeMap[row][col] {
			case pipeVertical, pipeHorizontal, pipeBendNorthEast, pipeBendNorthWest, pipeBendSouthWest, pipeBendSouthEast, ground, startingPosition:
			default:
				return nil, fmt.Errorf("unknown symbol %q", pipeMap[row][col])
			}
		}
	}

	return pipeMap, nil
}

func findLoop(pipeMap [][]byte) ([]Position, error) {
	startingPosition, err := getStartPosition(pipeMap)
	if err != nil {
		return nil, fmt.Errorf("could not find starting position: %w", err)
	}

	seen := make([][]bool, len(pipeMap))
	for i := range seen {
		seen[i] = make([]bool, len(pipeMap[i]))
	}

	loop := []Position{startingPosition}
	seen[startingPosition.row][startingPosition.col] = true

	for {
		neighbors := getConnectedNeighbors(pipeMap, loop[len(loop)-1])
		if len(neighbors) != 2 {
			return nil, fmt.Errorf("stumbled upon position with %d neighbors: %#v", len(neighbors), loop[len(loop)-1])
		}

		for len(neighbors) > 0 && seen[neighbors[0].row][neighbors[0].col] {
			neighbors = neighbors[1:]
		}

		if len(neighbors) == 0 {
			break
		}

		loop = append(loop, neighbors[0])
		seen[neighbors[0].row][neighbors[0].col] = true
	}

	return loop, nil
}

func getStartPosition(pipeMap [][]byte) (Position, error) {
	for row := range pipeMap {
		for col := range pipeMap[row] {
			if pipeMap[row][col] == startingPosition {
				return Position{row, col}, nil
			}
		}
	}

	return Position{}, fmt.Errorf("no starting position found")
}

func getConnectedNeighbors(pipeMap [][]byte, pos Position) []Position {
	var neighbors []Position

	shape := pipeMap[pos.row][pos.col]

	switch shape {
	case startingPosition:
		neighbors = getStartingPositionConnectedNeighbors(pipeMap, pos)
	case pipeVertical:
		if pos.row > 0 {
			neighbors = append(neighbors, Position{pos.row - 1, pos.col})
		}
		if pos.row < len(pipeMap)-1 {
			neighbors = append(neighbors, Position{pos.row + 1, pos.col})
		}
	case pipeHorizontal:
		if pos.col > 0 {
			neighbors = append(neighbors, Position{pos.row, pos.col - 1})
		}
		if pos.col < len(pipeMap[pos.row])-1 {
			neighbors = append(neighbors, Position{pos.row, pos.col + 1})
		}
	case pipeBendNorthEast:
		if pos.row > 0 {
			neighbors = append(neighbors, Position{pos.row - 1, pos.col})
		}
		if pos.col < len(pipeMap[pos.row])-1 {
			neighbors = append(neighbors, Position{pos.row, pos.col + 1})
		}
	case pipeBendNorthWest:
		if pos.row > 0 {
			neighbors = append(neighbors, Position{pos.row - 1, pos.col})
		}
		if pos.col > 0 {
			neighbors = append(neighbors, Position{pos.row, pos.col - 1})
		}
	case pipeBendSouthWest:
		if pos.row < len(pipeMap)-1 {
			neighbors = append(neighbors, Position{pos.row + 1, pos.col})
		}
		if pos.col > 0 {
			neighbors = append(neighbors, Position{pos.row, pos.col - 1})
		}
	case pipeBendSouthEast:
		if pos.row < len(pipeMap)-1 {
			neighbors = append(neighbors, Position{pos.row + 1, pos.col})
		}
		if pos.col < len(pipeMap[pos.row])-1 {
			neighbors = append(neighbors, Position{pos.row, pos.col + 1})
		}
	}

	return neighbors
}

func getStartingPositionConnectedNeighbors(pipeMap [][]byte, pos Position) []Position {
	var neighbors []Position

	if pos.row > 0 && contains([]byte{pipeVertical, pipeBendSouthEast, pipeBendSouthWest}, pipeMap[pos.row-1][pos.col]) {
		neighbors = append(neighbors, Position{pos.row - 1, pos.col})
	}
	if pos.row < len(pipeMap)-1 && contains([]byte{pipeVertical, pipeBendNorthEast, pipeBendNorthWest}, pipeMap[pos.row+1][pos.col]) {
		neighbors = append(neighbors, Position{pos.row + 1, pos.col})
	}
	if pos.col > 0 && contains([]byte{pipeHorizontal, pipeBendNorthEast, pipeBendSouthEast}, pipeMap[pos.row][pos.col-1]) {
		neighbors = append(neighbors, Position{pos.row, pos.col - 1})
	}
	if pos.col < len(pipeMap[pos.row])-1 && contains([]byte{pipeHorizontal, pipeBendNorthWest, pipeBendSouthWest}, pipeMap[pos.row][pos.col+1]) {
		neighbors = append(neighbors, Position{pos.row, pos.col + 1})
	}

	return neighbors
}

func contains(values []byte, value byte) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
