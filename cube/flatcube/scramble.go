package flatcube

import "rubik/cube/moves"

func (_c *FlatCube) Scramble() {
	_c.MakeMoves(moves.GetRandom(10000))
}
