package twentyone

import (
	"aoc/help"
)

type Die struct {
	sides int
	next  int
	rolls int
}

func (d *Die) Roll() int {
	r := d.next + 1
	d.rolls += 1
	d.next = (d.next + 1) % d.sides
	return r
}

func (d *Die) Rolls() int {
	return d.rolls
}

func NewDie(sides int) Die {
	return Die{sides: sides}
}

func PartOne(playerOne int, playerTwo int) int {
	die := NewDie(100)
	oneScore := 0
	twoScore := 0
	playerOne -= 1
	playerTwo -= 1
	for {
		oneRoll := die.Roll() + die.Roll() + die.Roll()
		playerOne = (playerOne + oneRoll) % 10
		oneScore += playerOne + 1
		if oneScore >= 1000 {
			break
		}
		twoRoll := die.Roll() + die.Roll() + die.Roll()
		playerTwo = (playerTwo + twoRoll) % 10
		twoScore += playerTwo + 1
		if twoScore >= 1000 {
			break
		}
	}
	return die.Rolls() * help.Min(oneScore, twoScore)
}

type Result struct {
	One int
	Two int
}
type State struct {
	OnePlace int
	OneScore int
	TwoPlace int
	TwoScore int
	Turn     bool
}

func (s *State) Next(roll int) State {
	next := *s
	if next.Turn {
		place := (s.OnePlace+roll-1)%10 + 1
		return State{
			OnePlace: place,
			OneScore: s.OneScore + place,
			TwoPlace: s.TwoPlace,
			TwoScore: s.TwoScore,
			Turn:     !s.Turn,
		}
	} else {
		place := (s.TwoPlace+roll-1)%10 + 1
		return State{
			OnePlace: s.OnePlace,
			OneScore: s.OneScore,
			TwoPlace: place,
			TwoScore: s.TwoScore + place,
			Turn:     !s.Turn,
		}
	}
}

func (r *Result) Add(or Result) {
	r.One += or.One
	r.Two += or.Two
}

func (r *Result) Times(i int) Result {
	return Result{r.One * i, r.Two * i}
}

func Play(one int, two int) *Result {
	return play(State{one, 0, two, 0, true})

}

var cache = make(map[State]*Result)

func play(s State) *Result {
	r := &Result{0, 0}
	if r, ok := cache[s]; ok {
		return r
	}
	if s.OneScore >= 21 {
		r = &Result{1, 0}
	} else if s.TwoScore >= 21 {
		r = &Result{0, 1}
	} else {
		r.Add(play(s.Next(3)).Times(1))
		r.Add(play(s.Next(4)).Times(3))
		r.Add(play(s.Next(5)).Times(6))
		r.Add(play(s.Next(6)).Times(7))
		r.Add(play(s.Next(7)).Times(6))
		r.Add(play(s.Next(8)).Times(3))
		r.Add(play(s.Next(9)).Times(1))
	}
	cache[s] = r
	return r
}
