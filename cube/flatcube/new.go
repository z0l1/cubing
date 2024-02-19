package flatcube

func makeSide(color Side) [3][3]Side {
	return [3][3]Side{
		{color, color, color},
		{color, color, color},
		{color, color, color},
	}
}

func New() *RubikCubeFlat {
	cube := [6][3][3]Side{
		makeSide(WhiteSide),
		makeSide(RedSide),
		makeSide(BlueSide),
		makeSide(OrangeSide),
		makeSide(GreenSide),
		makeSide(YellowSide),
	}

	return &RubikCubeFlat{
		sides: cube,
	}
}
