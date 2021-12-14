package thirteen

import (
	"aoc/help"
	"fmt"
	"strings"
)

type Place struct {
	X int
	Y int
}
type Fold Place

type Paper map[Place]bool

func (p Paper) Fold(f Fold) Paper {
	var m func(o Place) Place
	np := Paper{}
	if f.X == 0 {
		m = func(o Place) Place {
			if o.Y > f.Y {
				return Place{o.X, f.Y - (o.Y - f.Y)}
			} else {
				return o
			}
		}
	} else {
		m = func(o Place) Place {
			if o.X > f.X {
				return Place{f.X - (o.X - f.X), o.Y}
			} else {
				return o
			}
		}
	}
	for place := range p {
		np[m(place)] = true
	}
	return np
}
func (p Paper) Apply(fs []Fold) Paper {
	for _, f := range fs {
		p = p.Fold(f)
	}
	return p
}

func (p Paper) Read() {
	mx := 0
	my := 0
	for place := range p {
		if place.X > mx {
			mx = place.X
		}
		if place.Y > my {
			my = place.Y
		}
	}
	out := strings.Builder{}
	for j := 0; j <= my; j++ {
		for i := 0; i <= mx; i++ {
			if _, ok := p[Place{i, j}]; ok {
				out.WriteRune('#')
			} else {
				out.WriteRune(' ')
			}
		}
		out.WriteRune('\n')
	}
	fmt.Println(out.String())
}

func GetInput(path string) (Paper, []Fold) {
	paper := Paper{}
	folds := []Fold{}

	lines := help.ReadInput(path)
	cut := 0
	for i, line := range lines {
		if line == "" {
			cut = i + 1
			break
		}
		coords := strings.SplitN(line, ",", 2)
		paper[Place{help.Sinter(coords[0]), help.Sinter(coords[1])}] = true
	}
	for _, line := range lines[cut:] {
		fold := strings.Split(strings.Split(line, " ")[2], "=")
		if fold[0] == "x" {
			folds = append(folds, Fold{help.Sinter(fold[1]), 0})
		} else {
			folds = append(folds, Fold{0, help.Sinter(fold[1])})
		}
	}

	return paper, folds
}
