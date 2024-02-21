package moves

type Move string

const (
	WhiteC  Move = "w"
	WhiteCC Move = "w'"

	RedC  Move = "r"
	RedCC Move = "r'"

	BlueC  Move = "b"
	BlueCC Move = "b'"

	GreenC  Move = "g"
	GreenCC Move = "g'"

	OrangeC  Move = "o"
	OrangeCC Move = "o'"

	YellowC  Move = "y"
	YellowCC Move = "y'"
)

var moves = []Move{
	WhiteC, WhiteCC,
	RedC, RedCC,
	BlueC, BlueCC,
	GreenC, GreenCC,
	OrangeC, OrangeCC,
	YellowC, YellowCC,
}

var movesMap = map[string]Move{
	string(WhiteC):   WhiteC,
	string(WhiteCC):  WhiteCC,
	string(RedC):     RedC,
	string(RedCC):    RedCC,
	string(BlueC):    BlueC,
	string(BlueCC):   BlueCC,
	string(GreenC):   GreenC,
	string(GreenCC):  GreenCC,
	string(OrangeC):  OrangeC,
	string(OrangeCC): OrangeCC,
	string(YellowC):  YellowC,
	string(YellowCC): YellowCC,
}

func GetAll() []Move {
	return moves
}

func Parse(move string) (Move, bool) {
	for m := range moves {
		if string(m) == move {
		}
	}

	mv, ok := movesMap[move]
	return mv, ok
}

//(side cube.SideColor, cc bool, err error)
