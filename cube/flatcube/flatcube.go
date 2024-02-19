package flatcube

import "encoding/json"

type SideColor int

const (
	White  SideColor = 0
	Red    SideColor = 1
	Blue   SideColor = 2
	Orange SideColor = 3
	Green  SideColor = 4
	Yellow SideColor = 5
)

type Direction int

const (
	Top    Direction = 0
	Bottom Direction = 1
	Left   Direction = 2
	Right  Direction = 3
)

type SideConnection struct {
	side SideColor
	dir  Direction
}

type CubeSide [3][3]SideColor

type FlatCube [6]CubeSide

type RubikCubeFlat struct {
	sides FlatCube
}

func (rcf *RubikCubeFlat) String() string {
	bytes, err := json.Marshal(rcf.sides)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func (rcf *RubikCubeFlat) GetSides() *FlatCube {
	return &rcf.sides
}

func (rcf *RubikCubeFlat) GetSide(side SideColor) *CubeSide {
	return &rcf.sides[side]
}
