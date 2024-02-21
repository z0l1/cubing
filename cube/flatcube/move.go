package flatcube

import (
	"rubik/cube"
	"rubik/cube/moves"
	"strings"
)

func (_c *FlatCube) makeMove(side cube.SideColor, cc bool) {
	RotateFace(&_c.sides[side], cc)
	RotateNeighbours(&_c.sides, side, cc)
}

func (_c *FlatCube) MakeMove(_move moves.Move) {
	moveStr := string(_move)
	_, moveOk := moves.Parse(string(_move))
	if !moveOk {
		return
	}

	side := cube.SideColor(strings.Index("wrbogy", string(moveStr[0])))
	cc := len(moveStr) == 2

	RotateFace(&_c.sides[side], cc)
	RotateNeighbours(&_c.sides, side, cc)
}

func (_c *FlatCube) MakeMoves(moves []moves.Move) {
	for _, move := range moves {
		_c.MakeMove(move)
	}
}
