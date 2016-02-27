package counter

import "testing"

type testDatum struct {
	strs                  []string
	chars, letters, lines int
	name                  string
}

var testData = []testDatum{
	{[]string{""}, 0, 0, 0, "empty string"},
	{[]string{"one"}, 3, 3, 1, "one line"},
	{[]string{"2323·2005·7766·355"}, 18, 0, 1, "one line with unicode"},
	{[]string{"resumé"}, 6, 6, 1, "one line with unicode"},
	{[]string{"count the chars,\nletters,\nand lines"}, 35, 28, 3, "spaces"},
	{[]string{"count the ch", "ars,\nlett", "ers,\nand lines"}, 35, 28, 3,
		"spaces, chopped up"},
	{[]string{"one\ntwo"}, 7, 6, 2, "two lines"},
	{[]string{"one\n"}, 4, 3, 1, "two lines, no chars on second line"},
	{[]string{"one\ntwo\nthree"}, 13, 11, 3, "three lines"},
	{[]string{"one\n", "two\n", "three"}, 13, 11, 3, "three lines"},
	{[]string{"one\ntwo\n"}, 8, 6, 2, "two lines, empty last line"},
	{[]string{"\n\n"}, 2, 0, 2, "two empty lines"},
}

func TestValid(t *testing.T) {
	for i, test := range testData {
		c := makeCounter()
		for _, s := range test.strs {
			c.AddString(s)
		}
		subTest(t, "chars", i, c.Characters(), test.chars, test)
		subTest(t, "letters", i, c.Letters(), test.letters, test)
		subTest(t, "lines", i, c.Lines(), test.lines, test)
	}
}

func subTest(t *testing.T, tstr string, i, computed, wanted int, test testDatum) {
	if computed != wanted {
		t.Fatalf("Counter(%d), %s = %d, want %d (%s)\n%v",
			i, tstr, computed, wanted, test.name, test.strs)
	}
}
