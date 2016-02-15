package foodchain

import "fmt"

const TestVersion = 1

type verse struct {
	animal string
	secondLine string
	caught string
	lastLine string
	lastVerse bool
}

var verses = []verse{
	{"fly", "", "", "", false},
	{"spider", "It wriggled and jiggled and tickled inside her.",
		" that wriggled and jiggled and tickled inside her", "", false},
	{"bird", "How absurd to swallow a bird!", "", "", false},
	{"cat", "Imagine that, to swallow a cat!", "", "", false},
	{"dog", "What a hog, to swallow a dog!", "", "", false},
	{"goat", "Just opened her throat and swallowed a goat!", "", "", false},
	{"cow", "I don't know how she swallowed a cow!", "", "", false},
	{"horse", "She's dead, of course!", "", "", true},
}

func Verse(n int) (s string) {
	n--
	s += fmt.Sprintf("I know an old lady who swallowed a %s.", verses[n].animal)
	if verses[n].secondLine != "" {
		s += fmt.Sprintf("\n%s", verses[n].secondLine)
	}
	if !verses[n].lastVerse {
		for i := n-1; i >= 0; i-- {
			s += fmt.Sprintf("\nShe swallowed the %s to catch the %s",
				verses[i+1].animal, verses[i].animal)
			if verses[i].caught != "" {
				s += verses[i].caught
			}
			s += "."
		}
		s += "\nI don't know why she swallowed the fly. Perhaps she'll die."
	}
	return
}

func Verses(ns ...int) (s string) {
	for i, n := range ns {
		if i > 0 {
			s += "\n\n"
		}
		s += Verse(n)
	}
	return
}

func Song() string {
	// I don't yet know how to turn len(verses) into a list of int values to
	// use as varargs for Verses.
	return Verses(1, 2, 3, 4, 5, 6, 7, 8)
}
