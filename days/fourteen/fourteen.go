package fourteen

import (
	"aoc/help"
	"fmt"
	"strings"
)

func GetInput(path string) Formula {
	lines := help.ReadInput(path)
	formula := Formula{Sequence: lines[0], Rules: make(map[string]string)}
	for _, line := range lines[2:] {
		parts := strings.SplitN(line, " -> ", 2)
		formula.Rules[parts[0]] = parts[1]
	}
	return formula
}

type Formula struct {
	Sequence string
	Rules    map[string]string
}

func (f *Formula) Iterate() int {
	count := 0
	out := strings.Builder{}
	for i := 0; i < len(f.Sequence); i++ {
		out.WriteRune(rune(f.Sequence[i]))
		if i < len(f.Sequence)-1 {
			ins := f.Rules[f.Sequence[i:i+2]]
			if len(ins) > 0 {
				count += 1
			}
			out.WriteString(ins)
		}
	}
	f.Sequence = out.String()
	return count
}

func (f Formula) Counts() map[rune]int {
	counts := make(map[rune]int)
	for _, c := range f.Sequence {
		counts[c] += 1
	}
	return counts

}
func (f Formula) MoLe() int {
	least := len(f.Sequence)
	most := 0
	for _, c := range f.Counts() {
		if c < least {
			least = c
		}
		if c > most {
			most = c
		}
	}
	return most - least
}

func PartOne(f Formula, iterations int) int {
	for i := 0; i < iterations; i++ {
		f.Iterate()
		fmt.Println(f.MoLe())
	}
	return f.MoLe()
}
