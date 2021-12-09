package nine

import (
	"aoc/help"
	"sort"
	"strings"
)

type Place struct {
	X int
	Y int
}

func (p Place) Plus(o Place) Place {
	return Place{p.X + o.X, p.Y + o.Y}
}

type Field map[Place]int

type Basin map[Place]bool

var Directions = []Place{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (f Field) Neighbors(p Place) []Place {
	out := []Place{}
	for _, d := range Directions {
		neighbor := p.Plus(d)
		_, ok := f[neighbor]
		if ok {
			out = append(out, neighbor)
		}
	}
	return out
}

func GetInput(path string) Field {
	f := Field{}
	for i, line := range help.ReadInput(path) {
		for j, s := range strings.Split(line, "") {
			f[Place{i, j}] = help.Sinter(s)
		}
	}
	return f
}

func (f Field) LowPlaces() []Place {
	low := []Place{}
	for p, v := range f {
		isLow := func() bool {
			for _, n := range f.Neighbors(p) {
				if f[n] <= v {
					return false
				}
			}
			return true
		}()
		if isLow {
			low = append(low, p)
		}
	}
	return low
}

func (f Field) DangerTotal() int {
	danger := 0
	for _, p := range f.LowPlaces() {
		danger += 1 + f[p]
	}
	return danger
}

// pop removes and returns one Place from the given map.
func pop(mp *map[Place]bool) Place {
	for p := range *mp {
		delete(*mp, p)
		return p
	}
	return Place{}
}

func (f Field) BasinFor(p Place) Basin {
	basin := Basin{}
	todo := map[Place]bool{p: true}
	for len(todo) > 0 {
		place := pop(&todo)
		if f[place] == 9 {
			continue
		}
		basin[place] = true
		for _, n := range f.Neighbors(place) {
			if _, ok := basin[n]; !ok {
				todo[n] = true
			}
		}
	}
	return basin
}

func (f Field) Basins() []Basin {
	basins := []Basin{}
	seen := map[Place]bool{}
	for p, v := range f {
		if v == 9 {
			continue
		}
		if _, ok := seen[p]; ok {
			continue
		}
		basin := f.BasinFor(p)
		basins = append(basins, basin)
		for bp := range basin {
			seen[bp] = true
		}
	}
	return basins

}

func (f Field) BasinScore() int {
	basins := f.Basins()
	sort.Slice(basins, func(i int, j int) bool {
		return len(basins[i]) > len(basins[j])
	})
	return len(basins[0]) * len(basins[1]) * len(basins[2])
}