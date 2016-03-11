package connect

func ResultOf(lines []string) (string, error) {
	// O wins if it connects top to bottom
	// X wins if it connects left to right
	board := stringsToBytes(lines)
	if connectedTopToBottom(board, 'O') {
		return "white", nil
	} else if connectedLeftToRight(board, 'X') {
		return "black", nil
	}
	return "", nil
}

func connectedTopToBottom(board [][]byte, player byte) bool {
	for col, b := range board[0] {
		if b == player {
			if pathFound(board, player, 0, col, func(row, col int) bool { return row == len(board)-1 }) {
				return true
			}
		}
	}
	return false
}

func connectedLeftToRight(board [][]byte, player byte) bool {
	for row, bytes := range board {
		if bytes[0] == player {
			if pathFound(board, player, row, 0, func(row, col int) bool { return col == len(board[0])-1 }) {
				return true
			}
		}
	}
	return false
}

func stringsToBytes(strings []string) [][]byte {
	bytes := make([][]byte, len(strings))
	for i, s := range strings {
		bytes[i] = []byte(s)
	}
	return bytes
}

func pathFound(board [][]byte, player byte, row, col int, success func(int, int) bool) bool {
	todo := [][2]int{[2]int{row, col}}
	for len(todo) != 0 {
		newTodo := [][2]int{}
		for _, loc := range todo {
			if success(loc[0], loc[1]) {
				return true
			}
			board[loc[0]][loc[1]] = '.' // don't re-visit
			for r := row - 1; r <= row+1; r++ {
				for c := col - 1; c <= col+1; c++ {
					if r >= 0 && r < len(board) &&
						c >= 0 && c < len(board[0]) &&
						!(r == row && c == c) &&
						board[r][c] == player {
						newTodo = append(newTodo, [2]int{r, c})
					}
				}
			}
		}
		todo = newTodo
	}
	return false
}
