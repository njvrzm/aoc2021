package fifteen

import (
	"aoc/help"
	"fmt"
	"strings"
)

// Place needs to be centralized one of these days...
type Place struct {
	X int
	Y int
}

func (p Place) Plus(o Place) Place {
	return Place{p.X + o.X, p.Y + o.Y}
}

type Path []Place

var (
	Left   = Place{-1, 0}
	Right  = Place{1, 0}
	Down   = Place{0, 1}
	Up     = Place{0, -1}
	origin = Place{0, 0}
)

var Directions = []Place{Left, Right, Up, Down}

type RiskMap struct {
	Risk   map[Place]int
	Width  int
	Height int
	Target Place
}

var infinity = 1 << 32

type BucketMap struct {
	buckets map[int]map[Place]bool
	Index   int
}

func NewBucketMap() BucketMap {
	return BucketMap{make(map[int]map[Place]bool), 0}
}
func (bm *BucketMap) Empty() bool {
	bucket, ok := bm.buckets[bm.Index]
	return !ok || len(bucket) == 0
}
func (bm *BucketMap) Next() bool {
	for bm.Empty() {
		bm.Index++
	}
	return true
}

func (bm *BucketMap) Add(place Place, risk int) {
	_, ok := bm.buckets[risk]
	if !ok {
		bm.buckets[risk] = make(map[Place]bool)
	}
	bm.buckets[risk][place] = true
}

func (bm *BucketMap) Move(place Place, before int, after int) {
	delete(bm.buckets[before], place)
	bm.Add(place, after)
}

func (bm *BucketMap) Current() map[Place]bool {
	return bm.buckets[bm.Index]
}

func Pop(bucket map[Place]bool) Place {
	for p := range bucket {
		delete(bucket, p)
		return p
	}
	// impossible
	return origin
}

func (rm RiskMap) DialItIn() int {
	buckets := NewBucketMap()
	buckets.Add(origin, 0)
	risk := map[Place]int{}
	for buckets.Next() {
		current := buckets.Current()
		for len(current) > 0 {
			place := Pop(current)
			if place == rm.Target {
				return buckets.Index
			}
			risk[place] = buckets.Index
			for _, neighbor := range rm.Neighbors(place) {
				tentative, ok := risk[neighbor]
				potential := rm.Risk[neighbor] + buckets.Index
				if !ok {
					buckets.Add(neighbor, potential)
				} else if potential < tentative {
					buckets.Move(neighbor, tentative, potential)
				}
			}
		}
	}
	return -1
}

func (rm RiskMap) Neighbors(place Place) []Place {
	neighbors := []Place{}
	for _, direction := range Directions {
		neighbor := place.Plus(direction)
		if _, ok := rm.Risk[neighbor]; ok {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (rm *RiskMap) FromLines(lines []string) {
	rm.Risk = make(map[Place]int)
	rm.Width = len(lines[0])
	rm.Height = len(lines)
	for i, line := range lines {
		for j, r := range line {
			rm.Risk[Place{i, j}] = help.Sinter(string(r))
		}
	}
	rm.Target = Place{rm.Width - 1, rm.Height - 1}
}

func (rm *RiskMap) Magnify() {
	newRisk := map[Place]int{}
	for place, risk := range rm.Risk {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				newRisk[Place{place.X + i*rm.Width, place.Y + j*rm.Height}] = (risk-1+i+j)%9 + 1
			}
		}
	}
	rm.Risk = newRisk
	rm.Width *= 5
	rm.Height *= 5
	rm.Target = Place{rm.Width - 1, rm.Height - 1}
}

func (rm *RiskMap) Show() {
	out := strings.Builder{}
	for i := 0; i < rm.Width; i++ {
		for j := 0; j < rm.Height; j++ {
			out.WriteString(fmt.Sprintf("%d", rm.Risk[Place{i, j}]))
		}
		out.WriteString("\n")
	}
	fmt.Println(out.String())
}

func GetInput(path string) RiskMap {
	rm := RiskMap{}
	rm.FromLines(help.ReadInput(path))
	return rm
}
