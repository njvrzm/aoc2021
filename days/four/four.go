package four

import (
	"aoc/help"
	"fmt"
	"strings"
	"unicode"
)

type Grid [5][5]int

type Board struct {
	Values Grid
	Called Grid
}

type Game struct {
	Boards []*Board
	Calls []int
	Place int
}

func GetInput(path string) Game {
	b := Game{}
	lines := help.ReadInput(path)
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	for _, s := range strings.Split(lines[0], ",") {
		b.Calls = append(b.Calls, help.Sinter(s))
	}
	for i := 2; i<len(lines); i += 6 {
		board := Board{}
		for j := 0; j < 5; j++ {
			for k, s := range strings.FieldsFunc(lines[i+j], unicode.IsSpace) {
				board.Values[j][k] = help.Sinter(s)
			}
		}
		b.Boards = append(b.Boards, &board)
	}
	return b
}

func (b *Board) Call(n int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.Values[i][j] == n {
				b.Called[i][j] = 1
				return true
			}
		}
	}
	return false
}

func (b Board) Winning() bool {
	for i := 0; i < 5; i++ {
		horizontal := 0
		vertical := 0
		for j := 0; j < 5; j++ {
			horizontal += b.Called[i][j]
			vertical += b.Called[j][i]
		}
		if horizontal == 5 || vertical == 5 {
			return true
		}
	}
	return false
}

func (b Board) Score() int {
	if !b.Winning() {
		return 0
	}
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.Called[i][j] == 0 {
				score += b.Values[i][j]
			}
		}
	}
	return score
}

func (b Board) Show() {
	out := strings.Builder{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var cell string
			if b.Called[i][j] == 1 {
				cell = "**"
			} else {
				cell = fmt.Sprintf("%02d", b.Values[i][j])
			}
			out.WriteString(cell + " ")
		}
		out.WriteString("\n")
	}
	fmt.Println(out.String())
}

func (g *Game) Call() int {
	if g.Place > len(g.Calls) {
		panic("Game overboard!")
	}
	call := g.Calls[g.Place]
	g.Place += 1

	for _, b := range g.Boards {
		b.Call(call)
		if b.Winning() {
			return call * b.Score()
		}
	}
	return 0
}

func (g *Game) CallMisere() int {
	if g.Place > len(g.Calls) {
		panic("Game overboard!")
	}
	call := g.Calls[g.Place]
	g.Place += 1

	called := 0
	score := 0
	for _, b := range g.Boards {
		if b.Winning() {
			continue
		}
		called += 1
		b.Call(call)
		if b.Winning() {
			score = call * b.Score()
		}
	}
	if called == 1 && score > 0 {
		return score
	}
	return 0
}
