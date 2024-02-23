package cubehandler

import (
	"rubik/cube"
	"rubik/cube/flatcube"
	"rubik/cube/moves"
	"time"
)

type HandlerState int

const (
	Idle HandlerState = iota
	Running
	Solved
	//Stopped
)

type CubeHandler interface {
	Start() *chan<- moves.Move
	Stop()

	IsSolved() bool

	// GetChannel() *chan<- moves.Move
}

type CubeHandlerImpl struct {
	cube      cube.MovableCube
	state     HandlerState
	startTime time.Time
	moveCh    *chan moves.Move
	moveCnt   int64
}

func (handler *CubeHandlerImpl) stopWithState(state HandlerState) {
	handler.state = state
	close(*handler.moveCh)
}

func (handler *CubeHandlerImpl) IsSolved() bool {
	return handler.state == Solved
}

func (handler *CubeHandlerImpl) Start() *chan<- moves.Move {
	if handler.state == Running {
		var sendOnly chan<- moves.Move = *handler.moveCh
		return &sendOnly
	}

	handler.cube.Scramble()

	// buffered for 10k moves
	ch := make(chan moves.Move, 10000)

	handler.startTime = time.Now()
	handler.moveCh = &ch
	handler.state = Running
	handler.moveCnt = 0

	go func() {
		for {
			select {
			case move := <-ch:
				handler.cube.MakeMove(move)
				handler.moveCnt++
				if handler.cube.IsSolved() {
					handler.stopWithState(Solved)
				}
			}
		}
	}()

	var sendOnly chan<- moves.Move = ch
	return &sendOnly
}

func (handler *CubeHandlerImpl) Stop() {
	handler.stopWithState(Idle)
}

func New() CubeHandler {
	return &CubeHandlerImpl{
		state:  Idle,
		moveCh: nil,
		cube:   flatcube.New(),
	}
}
