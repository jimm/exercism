package beer

import (
	"errors"
	"fmt"
	"strings"
)

func Verse(n int) (string, error) {
	if n > 99 || n < 0 {
		return "", errors.New("wha?")
	}

	if n == 0 {
		return `No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
`, nil
	}

	plural := "s"
	if n == 1 {
		plural = ""
	}
	oneOrIt := "one"
	if n == 1 {
		oneOrIt = "it"
	}
	oneLessPlural := "s"
	if n-1 == 1 {
		oneLessPlural = ""
	}
	oneLessStr := fmt.Sprintf("%d", n-1)
	if n-1 == 0 {
		oneLessStr = "no more"
	}
	return fmt.Sprintf(`%d bottle%s of beer on the wall, %d bottle%s of beer.
Take %s down and pass it around, %s bottle%s of beer on the wall.
`,
		n, plural, n, plural, oneOrIt, oneLessStr, oneLessPlural), nil
}

func Verses(upper, lower int) (string, error) {
	if upper < lower {
		return "", errors.New("reverse, repeat")
	}
	verses := []string{}
	for i := upper; i >= lower; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		verses = append(verses, verse)
	}
	return strings.Join(verses, "\n") + "\n", nil
}

func Song() string {
	song, _ := Verses(99, 0)
	return song
}
