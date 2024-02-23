package main

import (
	"fmt"
	"rubik/cube/moves"
	"rubik/cubehandler"
	"time"
)

func genMovesBenchmark(n int) []moves.Move {
	genStart := time.Now()
	_moves := moves.GetRandom(n)
	genElapsed := time.Since(genStart)
	fmt.Println("generating", n, "moves took", genElapsed)

	return _moves
}

func makeThreadedMoves(ch *chan<- moves.Move, threadCnt int) {
	n := 1000
	for threadI := range threadCnt {
		fmt.Println("starting thread", threadI)

		go func(i int) {
			fmt.Println("started generating moves")
			_moves := genMovesBenchmark(n)
			fmt.Println("finished generating moves", i)
			for _, _move := range _moves {
				fmt.Println("sending from thread", i)
				*ch <- _move
			}
		}(threadI)
	}
}

func main() {
	fmt.Println("hi")

	handler := cubehandler.New()
	moveCh := handler.Start()

	go makeThreadedMoves(moveCh, 5)

	for {
	}

	fmt.Println("bye")
}
