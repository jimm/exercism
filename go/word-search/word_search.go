package wordsearch

import (
	"errors"
	"fmt"
)

const testVersion = 2

type Dir struct {
	dx, dy int
}

var dirs = []Dir{
	Dir{1, 0},   // east
	Dir{1, -1},  // southeast
	Dir{0, -1},  // south
	Dir{-1, -1}, // southwest
	Dir{-1, 0},  // west
	Dir{-1, 1},  // northwest
	Dir{0, 1},   // north
	Dir{1, 1},   // northeast
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	var err error

	puzzleBytes := stringsToBytes(puzzle)
	answer := map[string][2][2]int{}
	for _, word := range words {
		wordBytes := []byte(word)
		answer[word], err = find(wordBytes, puzzleBytes)
		if err != nil {
			return nil, err
		}
	}
	return answer, nil
}

func find(word []byte, puzzle [][]byte) ([2][2]int, error) {
	for r, row := range puzzle {
		for c, ch := range row {
			if word[0] == byte(ch) {
				for _, dir := range dirs {
					if match(word, puzzle, r, c, dir) {
						wordLen := len(word)
						return [2][2]int{[2]int{c, r}, [2]int{c + (wordLen-1)*dir.dy,
								r + (wordLen-1)*dir.dx}},
							nil
					}
				}
			}
		}
	}
	return [2][2]int{[2]int{0, 0}, [2]int{0, 0}}, errors.New(fmt.Sprintf("not found: %s", string(word)))
}

func match(word []byte, puzzle [][]byte, r, c int, dir Dir) bool {
	for i := 0; i < len(word); i++ {
		if r < 0 || r >= len(puzzle[0]) || c < 0 || c >= len(puzzle) {
			return false
		}
		if puzzle[r][c] != word[i] {
			return false
		}
		r += dir.dx
		c += dir.dy
	}
	return true
}

func stringsToBytes(strings []string) [][]byte {
	bytes := make([][]byte, len(strings))
	for i, s := range strings {
		bytes[i] = []byte(s)
	}
	return bytes
}
