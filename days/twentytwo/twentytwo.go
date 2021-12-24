package twentytwo

import (
	"aoc/help"
	"sort"
	"strings"
)

type Span struct {
	Low  int
	High int
}

func (sp *Span) Size() int {
	return sp.High + 1 - sp.Low
}

func (sp *Span) Overlaps(osp Span) bool {
	return (sp.High >= osp.Low) && (osp.High >= sp.Low)
}

func (sp *Span) Intersection(osp Span) Span {
	points := []int{sp.Low, sp.High, osp.Low, osp.High}
	sort.Ints(points)
	return Span{points[1], points[2]}
}

type Cuboid struct {
	X  Span
	Y  Span
	Z  Span
	On bool
}

func (c *Cuboid) Volume() int {
	return c.X.Size() * c.Y.Size() * c.Z.Size()
}

func (c *Cuboid) Overlaps(o Cuboid) bool {
	return c.X.Overlaps(o.X) && c.Y.Overlaps(o.Y) && c.Z.Overlaps(o.Z)
}

func (c *Cuboid) Intersection(o Cuboid) Cuboid {
	return Cuboid{
		c.X.Intersection(o.X),
		c.Y.Intersection(o.Y),
		c.Z.Intersection(o.Z),
		true, // a bit hacky but necessary
	}
}

type Cuboids []Cuboid

// Union counts the total number of cells left on by the given sequence.
// The idea is this: go through the sequence backwards. This way we only
// need to consider each cuboid once in its role of contributing to the
// total. "off" cuboids contribute nothing; to count the contribution of
// an "on" cuboid:
// * Find its intersection with each later cuboid
// * Find the total volume of those intersections (hey! it's recursion!)
// * Subtract that from the cuboid's volume
func (c Cuboids) Union(localOnly bool) int {
	totalOn := 0
	top := len(c) - 1
	for i := top; i >= 0; i-- {
		cuboid := c[i]
		if localOnly && !cuboid.Local() {
			continue
		}
		if cuboid.On {
			intersections := Cuboids{}
			for j := i + 1; j <= top; j++ {
				if cuboid.Overlaps(c[j]) {
					intersections = append(intersections, cuboid.Intersection(c[j]))
				}
			}
			subtract := intersections.Union(localOnly)
			totalOn += cuboid.Volume() - subtract
		}
	}
	return totalOn
}

func (sp *Span) FromString(s string) {
	lh := strings.SplitN(s, "..", 2)
	sp.Low = help.Sinter(lh[0])
	sp.High = help.Sinter(lh[1])
}

func CuboidFromString(s string) Cuboid {
	c := Cuboid{}
	parts := strings.SplitN(s, " ", 2)
	c.On = parts[0] == "on"
	xyz := strings.SplitN(parts[1], ",", 3)
	c.X.FromString(strings.SplitN(xyz[0], "=", 2)[1])
	c.Y.FromString(strings.SplitN(xyz[1], "=", 2)[1])
	c.Z.FromString(strings.SplitN(xyz[2], "=", 2)[1])
	return c
}

func GetInput(path string) Cuboids {
	cuboids := Cuboids{}
	for _, line := range help.ReadInput(path) {
		cuboids = append(cuboids, CuboidFromString(line))
	}
	return cuboids
}

// The following functions only apply to part one

func (sp *Span) Local() bool {
	return sp.Low >= -50 && sp.High <= 50
}

func (c *Cuboid) Local() bool {
	return c.X.Local() && c.Y.Local() && c.Z.Local()
}
