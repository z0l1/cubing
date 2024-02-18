package flatcube

// could probably make sides indexed so its "fancier" to get neighbours
// but this explanation took as much time as making the function below

func getNeighbours(s Side) []Side {
	if s == WhiteSide {
		return []Side{BlueSide, OrangeSide, GreenSide, RedSide}
	}

	if s == RedSide {
		return []Side{WhiteSide, BlueSide, YellowSide, GreenSide}
	}

	if s == BlueSide {
		return []Side{WhiteSide, OrangeSide, YellowSide, RedSide}
	}

	if s == OrangeSide {
		return []Side{WhiteSide, GreenSide, YellowSide, BlueSide}
	}

	if s == GreenSide {
		return []Side{WhiteSide, RedSide, YellowSide, OrangeSide}
	}

	if s == YellowSide {
		return []Side{RedSide, BlueSide, OrangeSide, GreenSide}
	}

	return []Side{}
}
