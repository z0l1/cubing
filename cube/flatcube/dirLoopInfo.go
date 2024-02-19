package flatcube

func dirLoopInfo(_dir Direction) (start [2]int /*, end [2]int */, dir [2]int) {
	switch _dir {
	case Top:
		start = [2]int{0, 0}
		//end = [2]int{0, 2}
		dir = [2]int{0, 1}
	case Bottom:
		start = [2]int{2, 2}
		//end = [2]int{2, 0}
		dir = [2]int{0, -1}
	case Left:
		start = [2]int{2, 0}
		//end = [2]int{0, 0}
		dir = [2]int{-1, 0}
	case Right:
		start = [2]int{0, 2}
		//end = [2]int{2, 2}
		dir = [2]int{1, 0}
	}

	return
}
