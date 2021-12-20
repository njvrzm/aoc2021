package seventeen

import (
	"aoc/help"
	"strings"
)

// PartOne is a thought experiment. In the X direction we can hit
// any triangular number within the box and "stop" there; this lets
// us simply determine the maximum initial vy. Whatever init_vy is,
// after 2vy + 1 steps we'll return to 0 with vy for the next step
// being -(init_vy + 1). To maximize init_vy we want that to be
// equal to the bottom of the box.
func PartOne(b Box) int {
	init_vy := -b.Down - 1
	return init_vy * (init_vy + 1) / 2
}

// PartTwo counts the total number of initial velocities that hit
// any point within the box. This is less amenable to an analytic
// solution for complicated reasons ^_^
func PartTwo(b Box) int {
	hits := 0
	for xv := 1; xv <= b.Right; xv++ {
		for yv := b.Down; yv <= -b.Down; yv++ {
			if b.Hits(xv, yv) {
				hits += 1
			}
		}
	}
	return hits
}

func (b *Box) Hits(xv int, yv int) bool {
	x := 0
	y := 0
	for {
		x += xv
		y += yv
		yv -= 1
		if xv > 0 {
			xv -= 1
		} else if xv < 0 {
			xv += 1
		}
		if x > b.Right || y < b.Down {
			return false
		}
		if b.Left <= x && x <= b.Right && b.Down <= y && y <= b.Up {
			return true
		}
	}
}

type Box struct {
	Left  int
	Right int
	Up    int
	Down  int
}

func (b *Box) FromString(line string) {
	parts := strings.SplitN(line, ", ", 2)
	xPart, yPart := parts[0], parts[1]
	xParts := help.Sinters(strings.SplitN(xPart[2:], "..", 2))
	yParts := help.Sinters(strings.SplitN(yPart[2:], "..", 2))
	b.Left, b.Right = help.MinMax(xParts)
	b.Down, b.Up = help.MinMax(yParts)
}

func GetInput(path string) (b Box) {
	b.FromString(help.ReadInput(path)[0][13:])
	return
}
