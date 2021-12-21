package twenty

import (
	"aoc/help"
	"fmt"
	"strings"
)

type Locus struct {
	X int
	Y int
}

type World struct {
	Cells map[Locus]bool
	Dirty map[Locus]bool

	Rules      []bool
	generation int
}

func NewWorld() World {
	return World{
		Cells: make(map[Locus]bool),
		Dirty: make(map[Locus]bool),
		Rules: []bool{},
	}
}

func (w *World) Set(l Locus, b bool) {
	for _, n := range w.Neighbors(l) {
		w.Dirty[n] = true
	}
	w.Cells[l] = b
}

func (w *World) Get(l Locus) bool {
	v, ok := w.Cells[l]
	if !ok && w.Rules[0] {
		return w.generation%2 == 1
	}
	return v
}

func (w *World) Tick() {
	dirty := make([]Locus, 0, len(w.Dirty))
	for l := range w.Dirty {
		dirty = append(dirty, l)
	}
	w.Dirty = make(map[Locus]bool)
	next := map[Locus]bool{}
	for _, l := range dirty {
		key := 0
		for i, n := range w.Neighbors(l) {
			if w.Get(n) {
				key += 1 << (8 - i)
			}
			next[l] = w.Rules[key]
		}
	}
	for l, v := range next {
		w.Set(l, v)
	}
	w.generation += 1
}

func (w *World) Neighbors(l Locus) []Locus {
	n := []Locus{}
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			n = append(n, Locus{l.X + dx, l.Y + dy})
		}
	}
	return n
}

func (w *World) Count() int {
	count := 0
	for _, v := range w.Cells {
		if v {
			count++
		}
	}
	return count
}

func (w *World) Show() {
	lx := 0
	hx := 0
	ly := 0
	hy := 0
	for l := range w.Cells {
		if l.X < lx {
			lx = l.X
		}
		if l.X > hx {
			hx = l.X
		}
		if l.Y < ly {
			ly = l.Y
		}
		if l.Y > hy {
			hy = l.Y
		}
	}
	out := strings.Builder{}
	for y := ly - 1; y <= hy+1; y++ {
		for x := lx - 1; x <= hx+1; x++ {
			ch := '.'
			where := Locus{x, y}
			if w.Cells[where] && w.Dirty[where] {
				ch = '8'
			} else if w.Cells[where] {
				ch = '#'
			} else if w.Dirty[where] {
				ch = '_'
			}
			out.WriteRune(ch)
		}
		out.WriteRune('\n')
	}
	fmt.Println(out.String())
}

func GetInput(path string) World {
	w := NewWorld()
	lines := help.ReadInput(path)
	inRules := true
	offset := 0
	for i, line := range lines {
		if line == "" {
			inRules = false
			offset = i
			continue
		}
		if inRules {
			for _, ch := range line {
				w.Rules = append(w.Rules, ch == '#')
			}
		} else {
			for j, ch := range line {
				w.Set(Locus{j, i - offset}, ch == '#')
			}
		}
	}
	return w
}
