package flatcube

func New() *RubikCubeFlat {
	cube := make([]side, 6, 6)

	cube[WhiteSide] = makeSide("w")
	cube[RedSide] = makeSide("r")
	cube[BlueSide] = makeSide("b")
	cube[OrangeSide] = makeSide("o")
	cube[GreenSide] = makeSide("g")
	cube[YellowSide] = makeSide("y")

	return &RubikCubeFlat{
		cube,
	}
}
