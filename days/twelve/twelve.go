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

func (m Map) CountEgresses(source Cave, canRepeat bool, beenTo map[string]int) int {
	if source.Name == "end" {
		return 1
	}
	egresses := 0
	if strings.ToLower(source.Name) == source.Name {
		beenTo[source.Name] += 1
	}
	for name := range source.Exits {
		if name == "start" {
			continue
		}
		if beenTo[name] > 0 {
			if canRepeat {
				egresses += m.CountEgresses(m.GetCave(name), false, beenTo)
			}
		} else {
			egresses += m.CountEgresses(m.GetCave(name), canRepeat, beenTo)
		}
	}
	if strings.ToLower(source.Name) == source.Name {
		beenTo[source.Name] -= 1
	}
	return egresses
}
