package house

import "fmt"

var jacksStuff = []string{
	"the horse and the hound and the horn\nthat belonged to",
	"the farmer sowing his corn\nthat kept",
	"the rooster that crowed in the morn\nthat woke",
	"the priest all shaven and shorn\nthat married",
	"the man all tattered and torn\nthat kissed",
	"the maiden all forlorn\nthat milked",
	"the cow with the crumpled horn\nthat tossed",
	"the dog\nthat worried",
	"the cat\nthat killed",
	"the rat\nthat ate",
	"the malt\nthat lay in",
}

func Embed(s1, s2 string) string {
	return fmt.Sprintf("%s %s", s1, s2)
}

func Verse(s1 string, s2 []string, s3 string) (s string) {
	s = "This is"
	for _, str := range s2 {
		s += " " + str
	}
	s += " " + s3
	return
}

func Song() (song string) {
	for i := len(jacksStuff); i >= 0; i-- {
		if i < len(jacksStuff) {
			song += "\n\n"
		}
		song += Verse("", jacksStuff[i:], "the house that Jack built.")
	}
	fmt.Println("***")
	fmt.Println(song)
	return
}
