package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	rows [][]int
}

func New(s string) (*Matrix, error) {
	rows := [][]int{}
	numCols := -1
	for _, rowString := range strings.Split(s, "\n") {
		row := []int{}
		fields := strings.Fields(rowString)
		if numCols == -1 {
			numCols = len(fields)
		} else if numCols != len(fields) {
			return nil, errors.New("rows must all have same number of cols")
		}
		for _, colString := range fields {
			cell, err := strconv.ParseInt(colString, 10, 32)
			if err != nil {
				return nil, err
			}
			row = append(row, int(cell))
		}
		rows = append(rows, row)
	}
	return &Matrix{rows}, nil
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(m.rows))
	for i, row := range m.rows {
		rows[i] = make([]int, len(row))
		copy(rows[i], row)
	}
	return rows
}

func (m *Matrix) Cols() [][]int {
	cols := make([][]int, len(m.rows[0]))
	for i := 0; i < len(m.rows[0]); i++ {
		cols[i] = []int{}
		for j := 0; j < len(m.rows); j++ {
			cols[i] = append(cols[i], m.rows[j][i])
		}
	}
	return cols
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m.rows) || col < 0 || col >= len(m.rows[0]) {
		return false
	}
	m.rows[row][col] = val
	return true
}
