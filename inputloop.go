package main

import (
	"bufio"
	"fmt"
	"os"
	"rubik/cube/flatcube"
	"rubik/cube/moves"
	"strings"
)

func getInput(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n') // reads input until the first occurrence of '\n'
	return strings.TrimRight(text, "\n")
}

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
