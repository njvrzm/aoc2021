package twentytwo

import (
	"aoc/help"
	"strings"
)

type Span struct {
	Low  int
	High int
}

func (sp *Span) Local() bool {
	return sp.Low >= -50 && sp.High <= 50
}

func (sp *Span) Width() int {
	return sp.High + 1 - sp.Low
}

func (sp *Span) FromString(s string) {
	lh := strings.SplitN(s, "..", 2)
	sp.Low = help.Sinter(lh[0])
	sp.High = help.Sinter(lh[1])
}

type Cube struct {
	X     Span
	Y     Span
	Z     Span
	On    bool
	Where Locus
	index int
}

func (c *Cube) Local() bool {
	return c.X.Local() && c.Y.Local() && c.Z.Local()
}

func (c *Cube) Init() {
	c.index = -1
	c.Next()
}

func (c *Cube) Volume() int {
	return c.X.Width() * c.Y.Width() * c.Z.Width()
}

func (c *Cube) Done() bool {
	return c.index == c.Volume()
}

func (c *Cube) Next() {
	c.index += 1
	c.Where.X = c.X.Low + c.index/c.Y.Width()/c.Z.Width()
	c.Where.Y = c.Y.Low + (c.index/c.Z.Width())%c.Y.Width()
	c.Where.Z = c.Z.Low + c.index%c.Z.Width()
}

func CubeFromString(s string) Cube {
	c := Cube{}
	parts := strings.SplitN(s, " ", 2)
	c.On = parts[0] == "on"
	xyz := strings.SplitN(parts[1], ",", 3)
	c.X.FromString(strings.SplitN(xyz[0], "=", 2)[1])
	c.Y.FromString(strings.SplitN(xyz[1], "=", 2)[1])
	c.Z.FromString(strings.SplitN(xyz[2], "=", 2)[1])
	return c
}

type Cubes []Cube

type Locus struct {
	X int
	Y int
	Z int
}

func (c Cubes) EvaluateSlow() int {
	on := map[Locus]bool{}
	for _, cube := range c {
		if !cube.Local() {
			continue
		}
		for cube.Init(); !cube.Done(); cube.Next() {
			if cube.On {
				on[cube.Where] = true
			} else {
				delete(on, cube.Where)
			}
		}
	}
	return len(on)
}

func GetInput(path string) Cubes {
	cubes := Cubes{}
	for _, line := range help.ReadInput(path) {
		cubes = append(cubes, CubeFromString(line))
	}
	return cubes
}
