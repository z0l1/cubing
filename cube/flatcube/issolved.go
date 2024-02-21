package flatcube

func (_c *FlatCube) IsSolved() bool {
	for _, side := range _c.sides {
		sideColor := side[1][1]
		for _, sideRow := range side {
			for _, piece := range sideRow {
				if piece != sideColor {
					return false
				}
			}
		}
	}

	return true
}
