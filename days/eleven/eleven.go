package eleven

import (
	"aoc/help"
	"strings"
)

type Place struct {
	X int
	Y int
}

func (p Place) Plus(o Place) Place {
	return Place{p.X + o.X, p.Y + o.Y}
}

type Field struct {
	Energy map[Place]int
}

type Basin map[Place]bool

var Directions = []Place{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
}

func (f *Field) Neighbors(p Place) []Place {
	out := []Place{}
	for _, d := range Directions {
		neighbor := p.Plus(d)
		_, ok := f.Energy[neighbor]
		if ok {
			out = append(out, neighbor)
		}
	}
	return out
}

func GetInput(path string) Field {
	f := Field{Energy: make(map[Place]int)}
	for i, line := range help.ReadInput(path) {
		for j, s := range strings.Split(line, "") {
			f.Energy[Place{i, j}] = help.Sinter(s)
		}
	}
	return f
}

func (f *Field) Pump(place Place) bool {
	f.Energy[place] += 1
	return f.Energy[place] >  9
}

func (f *Field) Step() int {
	flashes := 0
	flashers := map[Place]bool{}
	flashed := map[Place]bool{}
	for place := range f.Energy {
		if f.Pump(place) {
			flashers[place] = true
		}
	}
	for len(flashers) > 0 {
		this := pop(&flashers)
		flashes += 1
		flashed[this] = true
		for _, p := range f.Neighbors(this) {
			if f.Pump(p) {
				if _, ok := flashed[p]; !ok {
					flashers[p] = true
				}
			}
		}
	}
	for p := range flashed {
		f.Energy[p] = 0
	}
	return flashes
}

func (f *Field) ManyStep(n int) int {
	flashes := 0
	for i := 0; i < n; i ++ {
		flashes += f.Step()
	}
	return flashes
}

// Seek steps forward until the number of flashes (returned by
// Step()) is equal to the number of octopodes
func (f *Field) Seek() int {
	i := 0
	size := len(f.Energy)
	for {
		i += 1
		if f.Step() == size {
			return i
		}
	}
}

// pop removes and returns one Place from the given map.
func pop(mp *map[Place]bool) Place {
	for p := range *mp {
		delete(*mp, p)
		return p
	}
	return Place{}
}
