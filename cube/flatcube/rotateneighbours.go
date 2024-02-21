package flatcube

import "rubik/cube"

type SideRow [3]cube.SideColor

func getRow(side *cube.FlatSide, start [2]int, step [2]int) *SideRow {
	var result SideRow
	for i := 0; i < 3; i++ {
		result[i] = side[start[0]][start[1]]

		start[0] += step[0]
		start[1] += step[1]
	}

	return &result
}

func writeRow(side *cube.FlatSide, start [2]int, step [2]int, src *SideRow) {
	for i := 0; i < 3; i++ {
		side[start[0]][start[1]] = src[i]

		start[0] += step[0]
		start[1] += step[1]
	}
}

func rotateNeighbours(sides *cube.FlatSides, side cube.SideColor, cc bool) {
	conns := getConnections(side)

	startI := 0
	lastI := 3
	step := 1
	if cc {
		startI = 3
		lastI = 0
		step = -1
	}

	// save last
	lastConn := conns[lastI]
	lastStart, lastStep := getDirLoopInfo(conns[lastI].dir)
	swapper := getRow(&sides[lastConn.side], lastStart, lastStep)

	for i := startI; i != lastI+step; i += step {
		currConn := conns[i]
		currStart, currStep := getDirLoopInfo(currConn.dir)
		currSide := &sides[currConn.side]

		tmp := getRow(currSide, currStart, currStep)
		writeRow(currSide, currStart, currStep, swapper)
		swapper = tmp
	}

}
