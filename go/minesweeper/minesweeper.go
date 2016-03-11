package minesweeper

import "errors"

func (b Board) Count() error {
	if err := b.valid(); err != nil {
		return err
	}
	// len(*b)
	return nil
}

func (b Board) valid() error {
	rows := len(b) - 2
	cols := len(b[0]) - 2
	if !b.validShape(cols) {
		return errors.New("invalid shape")
	}
	if !b.validTopBorder(0, cols) || !b.validTopBorder(rows+1, cols) {
		return errors.New("invalid top or bottom border")
	}
	for i := 1; i <= rows; i++ {
		if !b.validEdgeBorder(i, cols) {
			return errors.New("invalid edge border")
		}
		if !b.validRow(i, cols) {
			return errors.New("invalid characters in row")
		}
	}
	for row := 1; row <= rows; row++ {
		for col := 1; col <= cols; col++ {
			if b[row][col] == ' ' {
				count := b.cellCount(row, col)
				if count > 0 {
					b[row][col] = byte('0' + count)
				}
			}
		}
	}
	return nil
}

func (b Board) cellCount(row, col int) int {
	count := 0
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if b[r][c] == '*' {
				count++
			}
		}
	}
	return count
}

func (b Board) validShape(cols int) bool {
	for _, row := range b {
		if len(row) != cols+2 {
			return false
		}
	}
	return true
}

func (b Board) validTopBorder(row, cols int) bool {
	if b[row][0] != '+' || b[row][cols+1] != '+' {
		return false
	}
	for i := 1; i < cols-1; i++ {
		if b[row][i] != '-' {
			return false
		}
	}
	return true
}

func (b Board) validEdgeBorder(row, cols int) bool {
	return b[row][0] == '|' && b[row][cols+1] == '|'
}

func (b Board) validRow(row, cols int) bool {
	bytes := b[row]
	for i := 1; i < cols+1; i++ {
		if bytes[i] != ' ' && bytes[i] != '*' {
			return false
		}
	}
	return true
}
