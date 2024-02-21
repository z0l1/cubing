package cube

import "rubik/cube/moves"

type MovableCube interface {
	MakeMove(move moves.Move)
	MakeMoves(moves []moves.Move)

	Scramble()
	IsSolved() bool
	//GetMoveCount() int
}

type SideColor int

type FlatSide [3][3]SideColor

type FlatSides [6]FlatSide
