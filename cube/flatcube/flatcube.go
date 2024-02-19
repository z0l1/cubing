package flatcube

type Side int

const (
	WhiteSide  Side = 0
	RedSide    Side = 1
	BlueSide   Side = 2
	OrangeSide Side = 3
	GreenSide  Side = 4
	YellowSide Side = 5
)

type Direction int

const (
	Top    Direction = 0
	Bottom Direction = 1
	Left   Direction = 2
	Right  Direction = 3
)

type SideConnection struct {
	side Side
	dir  Direction
}

type FlatCube [6][3][3]Side

type RubikCubeFlat struct {
	sides FlatCube
}
