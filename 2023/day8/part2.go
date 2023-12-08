package main

func part2(file string) int {
	instructions, desertMap := loadInput(file)
	startNodes := ghostStartNodes(desertMap)

	var pathLengths []int
	for _, start := range startNodes {
		pathLengths = append(pathLengths, stepsOnSingleGhostPath(desertMap, instructions, start))
	}

	if len(pathLengths) == 1 {
		return pathLengths[0]
	}

	parallelPathLength := pathLengths[0]
	for i := 1; i < len(pathLengths); i++ {
		parallelPathLength = lcm(parallelPathLength, pathLengths[i])
	}

	return parallelPathLength
}

func ghostStartNodes(desertMap DesertMap) []string {
	var nodes []string
	for id := range desertMap {
		if isGhostStartNode(id) {
			nodes = append(nodes, id)
		}
	}
	return nodes
}

func isGhostStartNode(id string) bool {
	return id[2] == 'A'
}

func isGhostFinishNode(id string) bool {
	return id[2] == 'Z'
}

func lcm(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}

	return abs(a*b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func stepsOnSingleGhostPath(desertMap DesertMap, instructions []string, start string) int {
	steps := 0
	current := start

	for !isGhostFinishNode(current) {
		currentNode := desertMap[current]
		instruction := instructions[steps%len(instructions)]
		switch instruction {
		case "L":
			current = currentNode[0]
		case "R":
			current = currentNode[1]
		}
		steps += 1
	}

	return steps
}
