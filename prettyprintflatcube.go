package main

import (
	"fmt"
	"rubik/cube"
	"rubik/cube/flatcube"
	"rubik/printer"
)

const SquareChar = "\u2588"
const DoubleSquareChar = SquareChar + SquareChar

func c(side cube.SideColor) string {
	switch side {
	case cube.White:
		return printer.White + DoubleSquareChar + printer.Reset
	case cube.Red:
		return printer.Red + DoubleSquareChar + printer.Reset
	case cube.Blue:
		return printer.Blue + DoubleSquareChar + printer.Reset
	case cube.Orange:
		return printer.Orange + DoubleSquareChar + printer.Reset
	case cube.Green:
		return printer.Green + DoubleSquareChar + printer.Reset
	case cube.Yellow:
		return printer.Yellow + DoubleSquareChar + printer.Reset
	}

	return ""
}

func printCube(_cube *flatcube.FlatCube) {
	sides := _cube.GetSides()
	w := sides[cube.White]
	r := sides[cube.Red]
	b := sides[cube.Blue]
	o := sides[cube.Orange]
	g := sides[cube.Green]
	y := sides[cube.Yellow]

	fmt.Println("          ", c(w[0][0]), c(w[0][1]), c(w[0][2]), "          ")
	fmt.Println("          ", c(w[1][0]), c(w[1][1]), c(w[1][2]), "          ")
	fmt.Println("          ", c(w[2][0]), c(w[2][1]), c(w[2][2]), "          ")
	fmt.Println()

	fmt.Println(c(g[0][0]), c(g[0][1]), c(g[0][2]), " ", c(r[0][0]), c(r[0][1]), c(r[0][2]), " ", c(b[0][0]), c(b[0][1]), c(b[0][2]))
	fmt.Println(c(g[1][0]), c(g[1][1]), c(g[1][2]), " ", c(r[1][0]), c(r[1][1]), c(r[1][2]), " ", c(b[1][0]), c(b[1][1]), c(b[1][2]))
	fmt.Println(c(g[2][0]), c(g[2][1]), c(g[2][2]), " ", c(r[2][0]), c(r[2][1]), c(r[2][2]), " ", c(b[2][0]), c(b[2][1]), c(b[2][2]))
	fmt.Println()

	fmt.Println("          ", c(y[0][0]), c(y[0][1]), c(y[0][2]), "          ")
	fmt.Println("          ", c(y[1][0]), c(y[1][1]), c(y[1][2]), "          ")
	fmt.Println("          ", c(y[2][0]), c(y[2][1]), c(y[2][2]), "          ")
	fmt.Println()

	fmt.Println("          ", c(o[0][0]), c(o[0][1]), c(o[0][2]), "          ")
	fmt.Println("          ", c(o[1][0]), c(o[1][1]), c(o[1][2]), "          ")
	fmt.Println("          ", c(o[2][0]), c(o[2][1]), c(o[2][2]), "          ")
	fmt.Println()

	fmt.Println("-----------------------------------------------------------------------------------------------")
}
