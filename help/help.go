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
	if err != nil {panic(err)}
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

func Min(ns []int) int {
	t := ns[0]
	for _, v := range ns[1:] {
		if v < t {
			t = v
		}
	}
	return t
}
func Max(ns []int) int {
	t := ns[0]
	for _, v := range ns[1:] {
		if v > t {
			t = v
		}
	}
	return t
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}