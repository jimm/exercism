package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const testVersion = 3

type team struct {
	name                string
	wins, losses, draws int
}

func Tally(in io.Reader, out io.Writer) error {
	var teams = map[string]*team{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()

		team1, team2, result, err := parse(line)
		if err != nil {
			return err
		}
		if result == "comment" {
			continue
		}

		var t1, t2 *team
		var ok bool
		if t1, ok = teams[team1]; !ok {
			t1 = &team{name: team1}
			teams[team1] = t1
		}
		if t2, ok = teams[team2]; !ok {
			t2 = &team{name: team2}
			teams[team2] = t2
		}
		switch result {
		case "win":
			t1.wins++
			t2.losses++
		case "loss":
			t1.losses++
			t2.wins++
		case "draw":
			t1.draws++
			t2.draws++
		default:
			return errors.New("bad result string")
		}
	}
	tallyHo(teams, out)
	return nil
}

func parse(line string) (string, string, string, error) {
	if line == "" || line[0] == '#' {
		return "", "", "comment", nil
	}

	fields := strings.Split(line, ";")
	if len(fields) != 3 {
		return "", "", "", errors.New("malformed line")
	}
	if fields[0] == "" {
		return "", "", "comment", nil
	}
	return fields[0], fields[1], fields[2], nil
}

//**************** team functions ****************

func (t *team) gamesPlayed() int { return t.wins + t.draws + t.losses }

func (t *team) points() int { return t.wins*3 + t.draws }

//**************** sorting and output ****************

type byWinners []*team

func (w byWinners) Len() int { return len(w) }

func (w byWinners) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

func (w byWinners) Less(i, j int) bool {
	if w[j].points() < w[i].points() {
		return true
	} else if w[j].points() < w[i].points() {
		return false
	}
	if w[j].wins < w[i].wins {
		return true
	} else if w[j].wins > w[i].wins {
		return false
	}
	if w[j].name > w[i].name {
		return true
	}
	return false
}

func tallyHo(teams map[string]*team, out io.Writer) error {
	teamPointers := []*team{}
	for _, pt := range teams {
		teamPointers = append(teamPointers, pt)
	}
	sort.Sort(byWinners(teamPointers))

	out.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, team := range teamPointers {
		s := fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", team.name,
			team.gamesPlayed(),
			team.wins, team.draws, team.losses,
			team.points())
		out.Write([]byte(s))
	}
	return nil
}
