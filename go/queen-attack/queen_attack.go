package queenattack

import (
	"errors"
	"math"
)

func CanQueenAttack(white, black string) (bool, error) {
	whiteRow, whiteCol, whiteErr := coords(white)
	if whiteErr != nil {
		return false, whiteErr
	}
	blackRow, blackCol, blackErr := coords(black)
	if blackErr != nil {
		return false, blackErr
	}

	rowDiff := int(math.Abs(float64(whiteRow) - float64(blackRow)))
	colDiff := int(math.Abs(float64(whiteCol) - float64(blackCol)))
	switch {
	case rowDiff == 0 && colDiff == 0:
		return false, errors.New("illegal parameters: both pieces are on the same square")
	case rowDiff == 0 || colDiff == 0:
		return true, nil
	default:
		return (rowDiff - colDiff) == 0, nil
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

	row, err = byteToNum(position[0], 'a', 'f', "bad rank")
	if err == nil {
		col, err = byteToNum(position[1], '1', '8', "bad file")
	}
	return
}

func byteToNum(c, low, high byte, errMsg string) (int, error) {
	if c >= byte(low) && c <= byte(high) {
		return int(c - byte(low)), nil
	}
	return 0, errors.New(errMsg)
}
