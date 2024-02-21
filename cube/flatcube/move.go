package flatcube

import (
	"rubik/cube"
	"strings"
)

func (_c *FlatCube) Move(move string) {
	if len(move) > 2 || len(move) < 1 {
		return
	}

	valid := true
	side := cube.White
	cc := strings.Contains(move, "'")
	if cc {
		move = strings.Trim(move, "'")
	}

	switch move {
	case "w":
		side = cube.White
		break

	case "r":
		side = cube.Red
		break

	case "b":
		side = cube.Blue
		break

	case "o":
		side = cube.Orange
		break

	case "g":
		side = cube.Green
		break

	case "y":
		side = cube.Yellow
		break

	default:
		valid = false
	}

	if !valid {
		return
	}

	RotateFace(&_c.sides[side], cc)
	RotateNeighbours(&_c.sides, side, cc)
}
