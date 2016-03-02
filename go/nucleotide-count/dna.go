package dna

import (
	"errors"
	"fmt"
)

type Histogram map[byte]int

func DNA(strand string) Histogram {
	h := map[byte]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, r := range strand {
		h[byte(r)]++
	}
	return h
}

func (h Histogram) Count(b byte) (int, error) {
	val, ok := h[b]
	if ok {
		return val, nil
	}
	return 0, errors.New(fmt.Sprintf("unknown nucleotide %v", b))
}

func (h Histogram) Counts() Histogram {
	return h
}
