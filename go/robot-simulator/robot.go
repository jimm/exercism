package robot

const (
	N = iota
	E
	S
	W
)

// **************** step 1 ****************

func Right() {
	Facing++
	if Facing > W {
		Facing = N
	}
}

func Left() {
	Facing--
	if Facing < N {
		Facing = W
	}
}

func Advance() {
	switch Facing {
	case N:
		Y++
	case E:
		X++
	case S:
		Y--
	case W:
		X--
	}
}

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case E:
		return "E"
	case S:
		return "S"
	case W:
		return "W"
	default:
		return "?"
	}
}

// **************** step 2 ****************

type Action string

func Robot(cmds chan Command, act chan Action) {
	for cmd := range cmds {
		switch cmd {
		case 'A':
			act <- "advance"
		case 'R':
			act <- "right"
		case 'L':
			act <- "left"
		}
	}
	close(act)
}

func Room(extent Rect, place DirAt, actions chan Action, rep chan DirAt) {
	// Initialze global data (boo, hiss)
	X = int(place.Pos.Easting)
	Y = int(place.Pos.Northing)
	Facing = place.Dir

	for action := range actions {
		switch action {
		case "advance":
			oldX := X
			oldY := Y
			Advance()
			if !legal(extent, X, Y) {
				X = oldX
				Y = oldY
			} else {
				place.Pos.Easting = RU(X)
				place.Pos.Northing = RU(Y)
			}
		case "right":
			Right()
			place.Dir = Facing
		case "left":
			Left()
			place.Dir = Facing
		}
	}
	rep <- place
	resetGlobalData()
}

func legal(extent Rect, x, y int) bool {
	return x >= int(extent.Min.Easting) && x <= int(extent.Max.Easting) &&
		y >= int(extent.Min.Northing) && y <= int(extent.Max.Northing)
}

// **************** step 3 ****************

type Action3 struct {
	name, cmd string
}

func Robot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		log <- "robots must have names"
		action <- Action3{"", "done"}
		return
	}

scriptRunner:
	for _, cmd := range script {
		switch cmd {
		case 'A':
			action <- Action3{name, "advance"}
		case 'R':
			action <- Action3{name, "right"}
		case 'L':
			action <- Action3{name, "left"}
		default:
			log <- "illegal command"
			break scriptRunner
		}
	}
	action <- Action3{name, "done"}
}

func Room3(extent Rect, robots []Place, actions chan Action3, report chan []Place, log chan string) {
	if !initialRobotCheck(extent, robots, log) {
		cleanup(robots, report)
		return
	}

	numActiveRobots := len(robots)
	for action := range actions {
		i := robotIndex(robots, action.name)
		if i == -1 {
			log <- "unknown robot trying to sneak in"
			cleanup(robots, report)
			return
		}
		place := &robots[i]
		X = place.px()
		Y = place.py()
		Facing = place.DirAt.Dir

		switch action.cmd {
		case "advance":
			Advance()
			if !legal(extent, X, Y) {
				log <- "robots can't walk through walls"
			} else if otherRobotAt(robots, action.name, X, Y) {
				log <- "two robots can't be at the same place at the same time"
			} else {
				place.DirAt.Pos.Easting = RU(X)
				place.DirAt.Pos.Northing = RU(Y)
			}
		case "right":
			Right()
			place.Dir = Facing
		case "left":
			Left()
			place.Dir = Facing
		case "done":
			numActiveRobots--
			if numActiveRobots == 0 {
				cleanup(robots, report)
				return
			}
		}
	}
}

func otherRobotAt(robots []Place, name string, x, y int) bool {
	for _, robot := range robots {
		if name != robot.Name && x == robot.px() && y == robot.py() {
			return true
		}
	}
	return false
}

func initialRobotCheck(extent Rect, robots []Place, log chan string) bool {
	for i, place := range robots {
		for j := i + 1; j < len(robots); j++ {
			if place.Name == robots[j].Name {
				log <- "duplicate robot name"
				return false
			}
		}
		if !legal(extent, place.px(), place.py()) {
			log <- "robot must be inside the room at all times"
		} else if otherRobotAt(robots, place.Name, place.px(), place.py()) {
			for j := i + 1; j < len(robots); j++ {
				if place.px() == robots[j].px() && place.py() == robots[j].py() {
					log <- "two robots can't be at the same place at the same time"
					return false
				}
			}
		}
	}
	return true
}

func robotIndex(robots []Place, name string) int {
	for i, robot := range robots {
		if robot.Name == name {
			return i
		}
	}
	return -1
}

func (p Place) px() int { return int(p.DirAt.Pos.Easting) }
func (p Place) py() int { return int(p.DirAt.Pos.Northing) }

func cleanup(robots []Place, report chan []Place) {
	report <- robots
	resetGlobalData()
}

// resetGlobalData resets global (ugh) data
func resetGlobalData() {
	X = 0
	Y = 0
	Facing = N
}
