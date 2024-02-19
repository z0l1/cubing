package flatcube

import "strings"

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
	if len(move) > 2 || len(move) < 1 {
		return
	}

	valid := true
	side := White
	cc := strings.Contains(move, "'")

	switch move {
	case "w":
		side = White
		break

	case "r":
		side = Red
		break

	case "b":
		side = Blue
		break

	case "o":
		side = Orange
		break

	case "g":
		side = Green
		break

	case "y":
		side = Yellow
		break

	default:
		valid = false
	}

	if !valid {
		return
	}

	RotateFace(&cube.sides[side], cc)
	RotateNeighbours(&cube.sides, side, cc)
}
