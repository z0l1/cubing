package main

import (
	"fmt"
	"rubik/cube/flatcube"
	"rubik/printer"
)

const SquareChar = "\u2588"
const DoubleSquareChar = SquareChar + SquareChar

func c(side flatcube.SideColor) string {
	switch side {
	case flatcube.White:
		return printer.White + DoubleSquareChar + printer.Reset
	case flatcube.Red:
		return printer.Red + DoubleSquareChar + printer.Reset
	case flatcube.Blue:
		return printer.Blue + DoubleSquareChar + printer.Reset
	case flatcube.Orange:
		return printer.Orange + DoubleSquareChar + printer.Reset
	case flatcube.Green:
		return printer.Green + DoubleSquareChar + printer.Reset
	case flatcube.Yellow:
		return printer.Yellow + DoubleSquareChar + printer.Reset
	}

	return ""
}

func printCube(sides flatcube.flatCubeSides) {
	w := sides[flatcube.White]
	r := sides[flatcube.Red]
	b := sides[flatcube.Blue]
	o := sides[flatcube.Orange]
	g := sides[flatcube.Green]
	y := sides[flatcube.Yellow]

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
