package fourteen

import (
	"aoc/help"
	"strings"
)

type Pair struct {
	Left  rune
	Right rune
}

func (p Pair) Insert(r rune) [2]Pair {
	return [2]Pair{{p.Left, r}, {r, p.Right}}
}

func GetInput(path string) Formula {
	lines := help.ReadInput(path)
	seq := lines[0]
	formula := Formula{
		LastLetter: rune(seq[len(seq)-1]),
		Rules:      make(map[Pair][2]Pair),
		PairCount:  make(PairCount),
	}
	for _, line := range lines[2:] {
		parts := strings.SplitN(line, " -> ", 2)
		from := Pair{rune(parts[0][0]), rune(parts[0][1])}
		to := rune(parts[1][0])
		formula.Rules[from] = from.Insert(to)
	}
	for i := 0; i < len(seq)-1; i++ {
		pair := Pair{rune(seq[i]), rune(seq[i+1])}
		formula.PairCount[pair] += 1

	}
	return formula
}

type PairCount map[Pair]int

type Formula struct {
	LastLetter rune
	PairCount  PairCount
	Rules      map[Pair][2]Pair
}

func (f *Formula) CountUp() {
	pc := PairCount{}
	for pair, count := range f.PairCount {
		for _, replacement := range f.Rules[pair] {
			pc[replacement] += count
		}
	}
	f.PairCount = pc
}

func (f Formula) LetterCount() map[rune]int {
	lc := make(map[rune]int)
	for pair, count := range f.PairCount {
		lc[pair.Left] += count
	}
	lc[f.LastLetter] += 1
	return lc
}

func MostMinusLeast(lc map[rune]int) int {
	most := 0
	least := 1 << 48
	for _, c := range lc {
		if c > most {
			most = c
		}
		if c < least {
			least = c
		}
	}
	return most - least
}

func PartOne(f Formula, iterations int) int {
	for i := 0; i < iterations; i++ {
		f.CountUp()
	}
	lc := f.LetterCount()
	return MostMinusLeast(lc)
}
