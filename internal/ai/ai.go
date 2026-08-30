package ai

func FindMove(board [9]string, ai string, human string) int {
	if pos := winningMove(board, ai); pos != 0 {
		return pos
	}
	if pos := winningMove(board, human); pos != 0 {
		return pos
	}
	if board[4] == "" {
		return 5
	}
	corners := [4]int{1, 3, 7, 9}
	for _, c := range corners {
		if board[c-1] == "" {
			return c
		}
	}
	sides := [4]int{2, 4, 6, 8}
	for _, s := range sides {
		if board[s-1] == "" {
			return s
		}
	}
	return 0
}

func winningMove(board [9]string, mark string) int {
	lines := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}
	for _, line := range lines {
		a, b, c := line[0], line[1], line[2]
		count := 0
		emptyPos := -1
		for _, idx := range []int{a, b, c} {
			if board[idx] == mark {
				count++
			} else if board[idx] == "" {
				emptyPos = idx
			}
		}
		if count == 2 && emptyPos != -1 {
			return emptyPos + 1
		}
	}
	return 0
}
