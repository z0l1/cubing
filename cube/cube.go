package cube

type RubikCube interface {
	MakeMove(move string)
	MakeMoves(moves []string)
}

type SideColor int

type FlatSide [3][3]SideColor

type FlatSides [6]FlatSide
