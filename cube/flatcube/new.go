package flatcube

import "rubik/cube"

func makeSide(color cube.SideColor) cube.FlatSide {
	return cube.FlatSide{
		{color, color, color},
		{color, color, color},
		{color, color, color},
	}
}

func New() *FlatCube {
	cube := cube.FlatSides{
		makeSide(cube.White),
		makeSide(cube.Red),
		makeSide(cube.Blue),
		makeSide(cube.Orange),
		makeSide(cube.Green),
		makeSide(cube.Yellow),
	}

	return &FlatCube{
		sides: cube,
	}
}
