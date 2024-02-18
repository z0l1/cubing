package flatcube

const (
	WhiteSide  = 0
	RedSide    = 1
	BlueSide   = 2
	OrangeSide = 3
	GreenSide  = 4
	YellowSide = 5
)

type side [][]string

type RubikCubeFlat struct {
	cube []side
}

func makeSide(color string) side {
	return side{
		{color, color, color},
		{color, color, color},
		{color, color, color},
	}
}
