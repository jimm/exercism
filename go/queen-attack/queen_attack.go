package queenattack

import (
	"errors"
	"math"
)

func CanQueenAttack(white, black string) (bool, error) {
	white_row, white_col, white_err := coords(white)
	if white_err != nil {
		return false, white_err
	}
	black_row, black_col, black_err := coords(black)
	if black_err != nil {
		return false, black_err
	}

	row_diff := int(math.Abs(float64(white_row) - float64(black_row)))
	col_diff := int(math.Abs(float64(white_col) - float64(black_col)))
	switch {
	case row_diff == 0 && col_diff == 0:
		return false, errors.New("illegal parameters: both pieces are on the same square")
	case row_diff == 0 || col_diff == 0:
		return true, nil
	default:
		return (row_diff - col_diff) == 0, nil
	}
}

func coords(position string) (row, col int, err error) {
	row = 0
	col = 0
	err = nil
	if len(position) != 2 {
		err = errors.New("illegal parameter: " + position)
		return
	}

	if position[0] >= byte('a') && position[0] <= byte('f') {
		row = int(position[0] - byte('a'))
	} else {
		err = errors.New("bad rank")
	}
	if position[1] >= byte('1') && position[1] <= byte('8') {
		col = int(position[1] - byte('1'))
	} else {
		err = errors.New("bad file")
	}
	return
}
