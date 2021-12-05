package five

import (
	"aoc/help"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p *Point) FromString(s string) {
	parts := strings.SplitN(s, ",", 2)
	p.X, p.Y = help.Sinter(parts[0]), help.Sinter(parts[1])
}

// Tocrement returns an int that is one closer to t than
// i is, unless i == t, in which case it returns t.
func Tocrement(i int, t int) int {
	if t > i {
		return i + 1
	} else if t < i {
		return i - 1
	} else {
		return t
	}
}

// StepTowards if given two points that are on a horizontal,
// vertical or diagonal (abs(slope)==1) line returns a point
// that is one step closer to o along that line than p is. If
// o and p are the same point, it returns that point. Otherwise
// the behavior is undefined. Or rather, it is defined but we
// don't care.
func (p Point) StepTowards(o Point) Point {
	return Point{X: Tocrement(p.X, o.X), Y: Tocrement(p.Y, o.Y)}
}

type Segment struct {
	Start Point
	End Point
}

func (seg *Segment) FromString(s string) {
	parts := strings.SplitN(s, " -> ", 2)
	seg.Start.FromString(parts[0])
	seg.End.FromString(parts[1])
}

func (seg Segment) IsDiagonal() bool {
	return seg.Start.X != seg.End.X && seg.Start.Y != seg.End.Y
}

func (seg Segment) Walk() []Point {
	at := Point{seg.Start.X, seg.Start.Y}
	points := []Point{at}
	for at.X != seg.End.X || at.Y != seg.End.Y {
		at = at.StepTowards(seg.End)
		points = append(points, at)
	}
	return points
}

func GetInput(path string) []Segment {
	lines := []Segment{}
	for _, line := range help.ReadInput(path)  {
		s := Segment{}
		s.FromString(line)
		lines = append(lines, s)
	}
	return lines
}

func Overlaps(segs []Segment, withDiagonal bool) int {
	count := make(map[Point]int)
	for _, seg := range segs {
		if seg.IsDiagonal() && !withDiagonal {
			continue
		}
		for _, point := range seg.Walk() {
			count[point] += 1
		}
	}
	overlaps := 0
	for _, c := range count {
		if c > 1 {
			overlaps += 1
		}
	}
	return overlaps
}