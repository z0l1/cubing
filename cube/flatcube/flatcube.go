package flatcube

import (
	"encoding/json"
	"rubik/cube"
)

type Direction int

const (
	Top    Direction = 0
	Bottom Direction = 1
	Left   Direction = 2
	Right  Direction = 3
)

type SideConnection struct {
	side cube.SideColor
	dir  Direction
}

type FlatCube struct {
	sides cube.FlatSides
}

func (_c *FlatCube) String() string {
	bytes, err := json.Marshal(_c.sides)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func (_c *FlatCube) GetSides() *cube.FlatSides {
	return &_c.sides
}

func (_c *FlatCube) GetSide(side cube.SideColor) *cube.FlatSide {
	return &_c.sides[side]
}
