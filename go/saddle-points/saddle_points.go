package matrix

import "math"

type Pair struct {
	row, col int
}

func (m *Matrix) Saddle() []Pair {
	cols := m.Cols()
	// for each col, find list of min value indexes
	colMinIndexes := [][]int{}
	for _, col := range cols {
		colMinIndexes = append(colMinIndexes, indexesOfMin(col))
	}

	// for each row, find list of max value indexes
	rowMaxIndexes := [][]int{}
	for _, row := range m.Rows() {
		rowMaxIndexes = append(rowMaxIndexes, indexesOfMax(row))
	}

	pairs := []Pair{}
	for row, rmis := range rowMaxIndexes {
		for _, col := range rmis {
			if includes(colMinIndexes[col], row) {
				pairs = append(pairs, Pair{row, col})
			}
		}
	}
	return pairs
}

func indexesOfMin(vec []int) []int {
	minVal := math.MaxInt32
	indexes := []int{}
	for i, val := range vec {
		if val < minVal {
			minVal = val
			indexes = []int{i}
		} else if val == minVal {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func indexesOfMax(vec []int) []int {
	maxVal := math.MinInt32
	indexes := []int{}
	for i, val := range vec {
		if val > maxVal {
			maxVal = val
			indexes = []int{i}
		} else if val == maxVal {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func includes(vec []int, val int) bool {
	for _, v := range vec {
		if v == val {
			return true
		}
	}
	return false
}
