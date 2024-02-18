package flatcube

type Side int

const (
	WhiteSide  Side = 0
	RedSide    Side = 1
	BlueSide   Side = 2
	OrangeSide Side = 3
	GreenSide  Side = 4
	YellowSide Side = 5
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
