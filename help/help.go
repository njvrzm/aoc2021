package help

import (
	"bufio"
	"io"
	"os"
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
