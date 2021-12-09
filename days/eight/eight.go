package eight

import (
	"aoc/help"
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Signals []string
	Outputs []string
}


func GetInput(path string) []Entry {
	entries := []Entry{}
	for _, l := range help.ReadInput(path) {
		parts := strings.SplitN(l, " | ", 2)
		signals := strings.Split(parts[0], " ")
		for i, s := range signals {
			chars := strings.Split(s, "")
			sort.Strings(chars)
			signals[i] = strings.Join(chars, "")
		}
		outputs := strings.Split(parts[1], " ")
		for i, s := range outputs {
			chars := strings.Split(s, "")
			sort.Strings(chars)
			outputs[i] = strings.Join(chars, "")
		}
		entries = append(entries, Entry{signals, outputs})
	}
	return entries
}

func PartOne(entries []Entry) int {
	t := 0
	for _, e := range entries {
		for _, o := range e.Outputs {
			l := len(o)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				t += 1
			}
		}
	}
	return t
}

type LightSet map[rune]bool
func NewLightSet(lights string) LightSet {
	out := LightSet{}
	for _, r := range lights {
		out[r] = true
	}
	return out
}

func (ls LightSet) Intersection(other LightSet) LightSet {
	out := LightSet{}
	for k := range ls {
		_, ok := other[k]
		if ok {
			out[k] = true
		}
	}
	return out
}

func (ls LightSet) Equals(other LightSet) bool {
	inter := ls.Intersection(other)
	return len(inter) == len(ls) && len(ls) == len(other)
}

// NumberLights maps each integer to the set of lights
// that are on in its proper representation
var NumberLights = map[int]LightSet {
	0:  NewLightSet("abcefg"),
	1:  NewLightSet("cf"),
	2:  NewLightSet("acdeg"),
	3:  NewLightSet("acdfg"),
	4:  NewLightSet("bcdf"),
	5:  NewLightSet("abdfg"),
	6:  NewLightSet("abdefg"),
	7:  NewLightSet("acf"),
	8:  NewLightSet("abcdefg"),
	9:  NewLightSet("abcdfg"),
}

// Signature returns the signature of the given string of
// lights
func Signature(ls LightSet, others []LightSet) string {
	sig := []string{}
	for _, other := range others {
		if ls.Equals(other) {
			continue
		}
		sig = append(sig, fmt.Sprint(len(ls.Intersection(other))))
	}
	sort.Strings(sig)
	return strings.Join(sig, "")
}

// SignatureToNumber returns a map of signature -> n, where
// the signature is the sorted string of counts of
// lights shared with each other number. Eg. if the
// lights for 0, 1, 2 were "ab", "abc", and "c", the
// signature for 0 would be "02", because 0 shares
// no lights with 2 and 2 lights with 1.
func SignatureToNumber() map[string]int {
	sigs := make(map[string]int)
	lsSlice := make([]LightSet, len(NumberLights))
	for i, ls := range NumberLights {
		lsSlice[i] = ls
	}
	for i, ls := range NumberLights {
		sigs[Signature(ls, lsSlice)] = i
	}
	return sigs
}

func Identify(lStrings []string) map[string]int {
	sign := SignatureToNumber()
	lSets := make(map[string]LightSet)
	for _, s := range lStrings {
		lSets[s] = NewLightSet(s)
	}
	lsSlice := []LightSet{}
	for _, ls := range lSets {
		lsSlice = append(lsSlice, ls)
	}

	out := make(map[string]int)
	for s, ls := range lSets {
		out[s] = sign[Signature(ls, lsSlice)]
	}
	return out
}