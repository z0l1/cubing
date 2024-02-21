package main

import (
	"bufio"
	"os"
	"rubik/cube/flatcube"
	"rubik/cube/moves"
)

func inputCubeLoop() {
	reader := bufio.NewReader(os.Stdin)
	_cube := flatcube.New()

	inp := ""
	for inp != "q" {
		_cube.MakeMove(moves.Move(inp))
		printCube(_cube)
		inp = getInput(reader, "Your moves: ")
	}
}
