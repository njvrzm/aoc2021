package sixteen

import (
	"fmt"
	"strconv"
	"strings"
)

type Packet struct {
	Version  int
	TypeID   int
	Children []Packet
	Value    int
	Span     int
}

func (p *Packet) Read(r *Reader, n int) string {
	p.Span += n
	return r.Read(n)
}

func (p *Packet) ReadInt(r *Reader, n int) int {
	p.Span += n
	return r.ReadInt(n)
}

func (p *Packet) FromReader(r *Reader) {
	if p.TypeID == 4 { // literal
		bits := []string{}
		for {
			chunk := p.Read(r, 5)
			bits = append(bits, chunk[1:])
			if strings.HasPrefix(chunk, "0") {
				break
			}
		}
		p.Value = BitsToInt(strings.Join(bits, ""))
		return
	} else { // operator
		ltid := p.ReadInt(r, 1)
		if ltid == 0 {
			length := p.ReadInt(r, 15)
			for length > 0 {
				child := r.Next()
				p.Span += child.Span
				length -= child.Span
				p.Children = append(p.Children, child)
			}
		} else {
			count := p.ReadInt(r, 11)
			for i := 0; i < count; i++ {
				child := r.Next()
				p.Span += child.Span
				p.Children = append(p.Children, child)
			}
		}
	}
}

func (p *Packet) VersionSum() int {
	v := p.Version
	for _, child := range p.Children {
		v += child.VersionSum()
	}
	return v
}

func (p *Packet) Evaluate() int {
	switch p.TypeID {
	case 0:
		return p.Sum()
	case 1:
		return p.Product()
	case 2:
		return p.Minimum()
	case 3:
		return p.Maximum()
	case 5:
		return p.GreaterThan()
	case 6:
		return p.LessThan()
	case 7:
		return p.Equal()
	default:
		return p.Value
	}
}

func (p *Packet) Sum() int {
	t := 0
	for _, child := range p.Children {
		t += child.Evaluate()
	}
	return t
}
func (p *Packet) Product() int {
	t := 1
	for _, child := range p.Children {
		t *= child.Evaluate()
	}
	return t
}
func (p *Packet) Minimum() int {
	m := -1
	for _, child := range p.Children {
		e := child.Evaluate()
		if m < 0 || e < m {
			m = e
		}
	}
	return m
}
func (p *Packet) Maximum() int {
	m := -1
	for _, child := range p.Children {
		e := child.Evaluate()
		if e > m {
			m = e
		}
	}
	return m
}
func (p *Packet) GreaterThan() int {
	if p.Children[0].Evaluate() > p.Children[1].Evaluate() {
		return 1
	} else {
		return 0
	}
}
func (p *Packet) LessThan() int {
	if p.Children[0].Evaluate() < p.Children[1].Evaluate() {
		return 1
	} else {
		return 0
	}
}
func (p *Packet) Equal() int {
	if p.Children[0].Evaluate() == p.Children[1].Evaluate() {
		return 1
	} else {
		return 0
	}
}

func (p *Packet) Repr(prefix string, infix string, suffix string) string {
	out := strings.Builder{}
	for i, c := range p.Children {
		if i == 0 {
			out.WriteString(prefix)
		} else {
			out.WriteString(infix)
		}
		out.WriteString(c.ToString())
	}
	out.WriteString(suffix)
	return out.String()
}
func (p *Packet) ToString() string {
	switch p.TypeID {
	case 0:
		return p.Repr("(", " + ", ")")
	case 1:
		return p.Repr("(", " * ", ")")
	case 2:
		return p.Repr("min(", ", ", ")")
	case 3:
		return p.Repr("max(", ", ", ")")
	case 5:
		return p.Repr("(", " > ", ")")
	case 6:
		return p.Repr("(", " < ", ")")
	case 7:
		return p.Repr("(", " == ", ")")
	default:
		return fmt.Sprintf("%d", p.Value)
	}

}

type Reader struct {
	bits  string
	index int
}

func NewReader(hex string) Reader {
	bits := strings.Builder{}
	for i := range hex {
		b, _ := strconv.ParseUint(hex[i:i+1], 16, 4)
		bits.WriteString(fmt.Sprintf("%04b", b))
	}
	return Reader{bits.String(), 0}
}

func (r *Reader) Read(n int) string {
	out := r.bits[r.index : r.index+n]
	r.index += n
	return out
}

func (r *Reader) ReadInt(n int) int {
	i, _ := strconv.ParseUint(r.Read(n), 2, 32)
	return int(i)
}

func (r *Reader) Next() Packet {
	p := Packet{}
	p.Version = p.ReadInt(r, 3)
	p.TypeID = p.ReadInt(r, 3)
	p.FromReader(r)
	return p
}

func BitsToInt(s string) int {
	i, _ := strconv.ParseUint(s, 2, 64)
	return int(i)
}
