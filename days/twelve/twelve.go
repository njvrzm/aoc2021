package twelve

import (
	"aoc/help"
	"strings"
)

func GetInput(path string) Map {
	caveMap := Map{make(map[string]Cave)}
	for _, line := range help.ReadInput(path) {
		caves := strings.SplitN(line, "-", 2)
		caveMap.Connect(caves[0], caves[1])
	}
	return caveMap
}

type Cave struct {
	Name    string
	Exits   map[string]bool
	Visited int
}

type Map struct {
	Caves map[string]Cave
}

func (m *Map) GetCave(name string) Cave {
	cave, ok := m.Caves[name]
	if !ok {
		cave = Cave{name, map[string]bool{}, 0}
		m.Caves[name] = cave
	}
	return cave
}

func (m *Map) Connect(oneName, twoName string) {
	one := m.GetCave(oneName)
	two := m.GetCave(twoName)
	one.Exits[two.Name] = true
	two.Exits[one.Name] = true
}

type Path []Cave

func (p Path) Occurences(name string) int {
	occ := 0
	for _, n := range p {
		if n.Name == name {
			occ += 1
		}
	}
	return occ
}

func (m *Map) CountEgresses(path Path, canRepeat bool) int {
	tail := path[len(path)-1]
	if tail.Name == "end" {
		return 1
	}
	egresses := 0
	for name := range tail.Exits {
		if name == "start" {
			continue
		}
		didRepeat := false
		if strings.ToLower(name) == name {
			occ := path.Occurences(name)
			if occ == 1 {
				if !canRepeat {
					continue
				}
				didRepeat = true
			} else if occ > 1 {
				continue
			}
		}
		egresses += m.CountEgresses(append(path, m.GetCave(name)), canRepeat && !didRepeat)
	}
	return egresses
}
