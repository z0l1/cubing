package cube

type RubikCube interface {
	MakeMove(move string)
	MakeMoves(moves []string)
}
