package flatcube

func rotateSide(sides *FlatCube, side SideColor, cc bool) {
	_sides := *sides

	// rotate face
	RotateFace(&_sides[side], cc)

	// rotate connecting items

	//startI := 0
	//endI := 3
	//dir := 1
	//if cc {
	//	startI = 3
	//	endI = 0
	//	dir = -1
	//}
	//
}

func rotateSideC(cube *RubikCubeFlat, side SideColor) {
}

func rotateSideCC(cube *RubikCubeFlat, side SideColor) {
}

func (cube *RubikCubeFlat) Move(move string) {
	switch move {
	case "w":

	}
}
