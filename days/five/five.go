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

type Segment struct {
	Start Point
	End Point
}

func (seg *Segment) FromString(s string) {
	parts := strings.SplitN(s, " -> ", 2)
	seg.Start.FromString(parts[0])
	seg.End.FromString(parts[1])
}

func (seg Segment) IsHorizontal() bool {
	return seg.Start.Y == seg.End.Y
}
func (seg Segment) IsVertical() bool {
	return seg.Start.X == seg.End.X
}
func (seg Segment) IsDiagonal() bool {
	return !(seg.IsHorizontal() || seg.IsVertical())
}

func (seg Segment) Points(withDiagonal bool) []Point {
	// diagonals ignored
	points := []Point{}
	if seg.IsHorizontal() {
		step := 1
		delta := seg.Start.X - seg.End.X
		if delta < 0 {
			delta = -delta
			step = -step
		}
		for i := 0; i <= delta; i++ {
			points = append(points, Point{seg.Start.X - i * step, seg.Start.Y})
		}
	} else if seg.IsVertical() {
		step := 1
		delta := seg.Start.Y - seg.End.Y
		if delta < 0 {
			delta = -delta
			step = -step
		}
		for i := 0; i <= delta; i++ {
			points = append(points, Point{seg.Start.X, seg.Start.Y - i * step})
		}
	} else if withDiagonal {
		xStep := 1
		yStep := 1
		xDelta := seg.Start.X - seg.End.X
		yDelta := seg.Start.Y - seg.End.Y
		if xDelta < 0 {
			xDelta = -xDelta
			xStep = -xStep
		}
		if yDelta < 0 {
			yDelta = -yDelta
			yStep = -yStep
		}
		for i := 0; i <= xDelta; i++ {
			points = append(points, Point{seg.Start.X - i * xStep, seg.Start.Y - i * yStep})
		}

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
		for _, point := range seg.Points(withDiagonal) {
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