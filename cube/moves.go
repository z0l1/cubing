package cube

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

func GetMoves() []Move {
	return []Move{
		WhiteC, WhiteCC,
		RedC, RedCC,
		BlueC, BlueCC,
		GreenC, GreenCC,
		OrangeC, OrangeCC,
		YellowC, YellowCC,
	}
}
