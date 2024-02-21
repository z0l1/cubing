package main

import (
	"fmt"
	"math/rand/v2"
	"rubik/cube/flatcube"
	"rubik/cube/moves"
	"time"
)

func getRandomMoves(n int) []moves.Move {
	allMoves := moves.GetAll()
	m := len(allMoves) - 1
	moves := make([]moves.Move, n, n)

	for i := 0; i < n; i++ {
		moves[i] = allMoves[rand.IntN(m)]
	}

	return moves

}

func genMovesBenchmark(n int) []moves.Move {
	genStart := time.Now()
	moves := getRandomMoves(n)
	genElapsed := time.Since(genStart)
	fmt.Println("generating", n, "moves took", genElapsed)

	return moves
}

func main() {
	fmt.Println("hi")
	_cube := flatcube.New()

	m := 10
	n := 1000 * 1000
	_moves := genMovesBenchmark(n)

	for i := range m {
		moveStart := time.Now()
		_cube.MakeMoves(_moves)
		moveElapsed := time.Since(moveStart)
		fmt.Println(i, ".: moving cube", len(_moves), "times took", moveElapsed)
	}

	fmt.Println("bye")
}
