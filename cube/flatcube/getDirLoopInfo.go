package flatcube

func getDirLoopInfo(_dir Direction) (start [2]int /*, end [2]int */, step [2]int) {
	switch _dir {
	case Top:
		start = [2]int{0, 0}
		//end = [2]int{0, 2}
		step = [2]int{0, 1}
	case Bottom:
		start = [2]int{2, 2}
		//end = [2]int{2, 0}
		step = [2]int{0, -1}
	case Left:
		start = [2]int{2, 0}
		//end = [2]int{0, 0}
		step = [2]int{-1, 0}
	case Right:
		start = [2]int{0, 2}
		//end = [2]int{2, 2}
		step = [2]int{1, 0}
	}

	return
}
