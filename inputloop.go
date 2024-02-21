package main

import (
	"bufio"
	"os"
	"rubik/cube/flatcube"
)

func inputCubeLoop() {
	reader := bufio.NewReader(os.Stdin)
	_cube := flatcube.New()

	inp := ""
	for inp != "q" {
		_cube.Move(inp)
		printCube(*_cube.GetSides())
		inp = getInput(reader, "Your move: ")
	}
}
