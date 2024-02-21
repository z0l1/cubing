package flatcube

func makeSide(color SideColor) CubeSide {
	return CubeSide{
		{color, color, color},
		{color, color, color},
		{color, color, color},
	}
}

func New() *FlatCube {
	cube := [6]CubeSide{
		makeSide(White),
		makeSide(Red),
		makeSide(Blue),
		makeSide(Orange),
		makeSide(Green),
		makeSide(Yellow),
	}

	return &FlatCube{
		sides: cube,
	}
}
