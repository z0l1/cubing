package flatcube

var corners = [4][2]int{
	// top left
	{0, 0},
	// top right
	{0, 2},
	// bottom right
	{2, 2},
	// bottom left
	{2, 0},
}

var mids = [4][2]int{
	// top
	{0, 1},
	// right
	{1, 2},
	// bottom
	{2, 1},
	// left
	{1, 0},
}

func RotateFace(side *CubeSide, cc bool) {
	LEN := len(corners)

	startI := 0
	endI := 3
	dir := 1
	stopValue := LEN

	if cc {
		startI = LEN - 1
		endI = 0
		dir = -1
		stopValue = -1
	}

	endCorner := corners[endI]
	endMid := mids[endI]

	// save last
	cornerSwapper := side[endCorner[0]][endCorner[1]]
	midSwapper := side[endMid[0]][endMid[1]]

	// swap current with last saved
	for i := startI; i != stopValue; i += dir {
		currentCorner := corners[i]
		currentMid := mids[i]

		tmp := side[currentCorner[0]][currentCorner[1]]
		side[currentCorner[0]][currentCorner[1]] = cornerSwapper
		cornerSwapper = tmp

		tmp = side[currentMid[0]][currentMid[1]]
		side[currentMid[0]][currentMid[1]] = midSwapper
		midSwapper = tmp
	}
}
