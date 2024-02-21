package main

import (
	"bufio"
	"fmt"
	"os"
	"rubik/cube/flatcube"
)

func main() {
	cube := flatcube.New()
	reader := bufio.NewReader(os.Stdin)

	inp := ""
	for inp != "q" {
		cube.Move(inp)
		printCube(*cube.GetSides())
		inp = getInput(reader, "Your move: ")
	}

	fmt.Println("bye")
}
