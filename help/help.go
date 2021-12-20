package help

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadInput(path string) []string {
	fh, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	lines := []string{}
	sc := bufio.NewScanner(io.Reader(fh))
	for sc.Scan() {
		lines = append(lines, strings.Trim(sc.Text(), "\n"))
	}
	return lines

}

func Sinter(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Sinters(ss []string) []int {
	ints := []int{}
	for _, s := range ss {
		ints = append(ints, Sinter(s))
	}
	return ints
}

func Min(ns ...int) int {
	t := ns[0]
	for _, v := range ns[1:] {
		if v < t {
			t = v
		}
	}
	return t
}
func Max(ns ...int) int {
	t := ns[0]
	for _, v := range ns[1:] {
		if v > t {
			t = v
		}
	}
	return t
}

func PosiMod(n int, modulus int) int {
	m := n % modulus
	if m < 0 {
		m += modulus
	}
	return m
}

func MinMax(ns []int) (int, int) {
	if len(ns) == 0 {
		return 0, 0
	}
	min := ns[0]
	max := ns[0]
	for _, n := range ns[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
