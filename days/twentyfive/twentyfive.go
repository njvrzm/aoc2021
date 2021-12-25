package twentyfive

import (
	"fmt"
	"strings"
)

type Locus struct {
	X int
	Y int
}

type World struct {
	Things map[Locus]rune
	Steps  int
	Width  int
	Height int
}

func (w *World) FromLines(lines []string) {
	w.Things = make(map[Locus]rune)
	w.Height = len(lines)
	w.Width = len(lines[0])
	for j, line := range lines {
		for i, r := range []rune(line) {
			switch r {
			case '>', 'v':
				w.Things[Locus{i, j}] = r
			}
		}
	}
}

func (w *World) Right(l Locus) Locus {
	return Locus{(l.X + 1) % w.Width, l.Y}
}
func (w *World) Down(l Locus) Locus {
	return Locus{l.X, (l.Y + 1) % w.Height}
}

func (w *World) Open(l Locus) bool {
	_, ok := w.Things[l]
	return !ok
}
func (w *World) Step() bool {
	w.Steps += 1
	horizontalMoves := []Locus{}
	for l, r := range w.Things {
		if r == '>' {
			if w.Open(w.Right(l)) {
				horizontalMoves = append(horizontalMoves, l)
			}
		}
	}
	for _, l := range horizontalMoves {
		delete(w.Things, l)
		w.Things[w.Right(l)] = '>'
	}
	verticalMoves := []Locus{}
	for l, r := range w.Things {
		if r == 'v' {
			if w.Open(w.Down(l)) {
				verticalMoves = append(verticalMoves, l)
			}
		}
	}
	for _, l := range verticalMoves {
		delete(w.Things, l)
		w.Things[w.Down(l)] = 'v'
	}
	return len(horizontalMoves) > 0 || len(verticalMoves) > 0
}

func (w *World) Show() {
	out := strings.Builder{}
	for j := 0; j < w.Height; j++ {
		for i := 0; i < w.Width; i++ {
			r, ok := w.Things[Locus{i, j}]
			if !ok {
				r = '.'
			}
			out.WriteRune(r)
		}
		out.WriteRune('\n')
	}
	fmt.Println(out.String())
}
