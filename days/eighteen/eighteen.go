package eighteen

import (
	"aoc/help"
	//"aoc/help"
	"fmt"
	"unicode"
)

func ReadNumbers(nums []string) []*Number {
	numbers := []*Number{}
	for _, ns := range nums {
		numbers = append(numbers, NumberFromString(ns))
	}
	return numbers
}

type Number struct {
	Left  *Number
	Right *Number
	Value int
}

type Direction int

const (
	LEFT  Direction = -1
	NONE  Direction = 0
	RIGHT Direction = 1
	UP    Direction = 2
)

func (n *Number) Depth() int {
	if n.IsLeaf() {
		return 0
	}
	l := n.Left.Depth()
	r := n.Right.Depth()
	if l > r {
		return 1 + l
	} else {
		return 1 + r
	}
}
func (n *Number) Truncate() (*Number, *Number) {
	l := n.Left
	r := n.Right
	n.Left = nil
	n.Right = nil
	return l, r
}

func (n *Number) IsLeaf() bool {
	// todo: add some validation in the construction process?
	return n.Left == nil && n.Right == nil
}

func (n *Number) Copy() *Number {
	if n.IsLeaf() {
		return &Number{Value: n.Value}
	} else {
		return &Number{Left: n.Left.Copy(), Right: n.Right.Copy()}
	}
}

func (n *Number) ToString() string {
	if n.IsLeaf() {
		return fmt.Sprintf("%d", n.Value)
	}
	return "[" + n.Left.ToString() + "," + n.Right.ToString() + "]"
}

func (n *Number) Plus(o *Number) *Number {
	if n == nil {
		return o.Copy()
	}
	sum := Number{Left: n.Copy(), Right: o.Copy()}
	sum.Reduce()
	return &sum
}

func (n *Number) Leftmost() *Number {
	node := n
	for !node.IsLeaf() {
		node = node.Left
	}
	return node
}
func (n *Number) Rightmost() *Number {
	node := n
	for !node.IsLeaf() {
		node = node.Right
	}
	return node
}

func (n *Number) Reduce() {
	for n.Explode() || n.Split() {
	}
}

func (n *Number) Explode() bool {
	_, d := n.explode(1)
	return d != NONE
}

func (n *Number) explode(depth int) (int, Direction) {
	if n.IsLeaf() {
		return 0, NONE
	}
	if depth == 4 {
		if !n.Left.IsLeaf() {
			l, r := n.Left.Truncate()
			if !l.IsLeaf() || !r.IsLeaf() {
				panic(fmt.Sprintf("Truncated non-leaf numbers: %v, %v", l, r))
			}
			n.Right.Leftmost().Value += r.Value
			return l.Value, LEFT
		} else if !n.Right.IsLeaf() {
			l, r := n.Right.Truncate()
			if !l.IsLeaf() || !r.IsLeaf() {
				panic(fmt.Sprintf("Truncated non-leaf numbers: %v, %v", l, r))
			}
			n.Left.Rightmost().Value += l.Value
			return r.Value, RIGHT
		} else {
			return 0, NONE
		}
	} else {
		v, d := n.Left.explode(depth + 1)
		switch d {
		case RIGHT:
			n.Right.Leftmost().Value += v
			return 0, UP
		case LEFT:
			return v, d
		case UP:
			return 0, UP
		case NONE:
			v, d = n.Right.explode(depth + 1)
			switch d {
			case LEFT:
				n.Left.Rightmost().Value += v
				return 0, UP
			case RIGHT:
				return v, d
			case UP:
				return 0, UP
			case NONE:
				return 0, NONE
			}
		}
	}
	return 0, UP
}

func (n *Number) Split() bool {
	if !n.IsLeaf() {
		return n.Left.Split() || n.Right.Split()
	} else if n.Value > 9 {
		n.Left = &Number{Value: n.Value / 2}
		n.Right = &Number{Value: n.Value - n.Left.Value}
		n.Value = 0
		return true
	}
	return false
}

func (n *Number) Magnitude() int {
	if n.IsLeaf() {
		return n.Value
	} else {
		return 3*n.Left.Magnitude() + 2*n.Right.Magnitude()
	}
}

func Sum(numbers []*Number) (sum *Number) {
	for _, num := range numbers {
		sum = sum.Plus(num)
	}
	return sum
}

type Reader struct {
	string string
	index  int
}

func (r *Reader) Peek() rune {
	return rune(r.string[r.index])
}

func (r *Reader) Next() rune {
	for r.Peek() == ' ' {
		r.index += 1
	}
	r.index += 1
	return rune(r.string[r.index-1])
}

func (r *Reader) Require(e rune) {
	n := r.Next()
	if n != e {
		panic(fmt.Sprintf("needed %c; got %c", e, n))
	}
}

func (r *Reader) ReadNumber() *Number {
	c := r.Next()
	if c == '[' {
		left := r.ReadNumber()
		r.Require(',')
		right := r.ReadNumber()
		r.Require(']')
		return &Number{Left: left, Right: right}
	} else if unicode.IsDigit(c) {
		return &Number{Value: int(c) - 48}
	} else {
		panic(fmt.Sprintf("Unexpected character %c", c))
	}
}

func NumberFromString(s string) *Number {
	r := &Reader{string: s}
	return r.ReadNumber()
}

func MaximumBinarySum(ns []*Number) int {
	max := 0
	for i := range ns {
		for j := range ns {
			max = help.Max(max, ns[i].Plus(ns[j]).Magnitude(), ns[j].Plus(ns[i]).Magnitude())

		}
	}
	return max
}
