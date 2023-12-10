package main

func part1() (any, error) {

	pipeMap, err := loadInput()
	if err != nil {
		return nil, err
	}

	loop, err := findLoop(pipeMap)
	if err != nil {
		return nil, err
	}

	return len(loop) / 2, nil
}
