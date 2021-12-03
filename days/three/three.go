package three

import (
	"aoc/help"
	"strconv"
)

func GetInput(path string) []string {
	return help.ReadInput(path)
}

func PartOne(input []string) int {
	size := len(input[0])
	gamma, epsilon := make([]byte, size), make([]byte, size)

	for i := 0; i < size; i++ {
		gamma[i] = commonestBitAt(input, i)
		epsilon[i] = rarestBitAt(input, i)
	}
	return bitsToInt(string(gamma)) * bitsToInt(string(epsilon))
}

func PartTwo(input []string) int {
	oxygen := whittle(input, commonestBitAt)
	carbox := whittle(input, rarestBitAt)
	return bitsToInt(oxygen) * bitsToInt(carbox)
}

func drop(ss []string, dd []bool) []string {
	out := []string{}
	for i, s  := range ss {
		if !dd[i] {
			out = append(out, s)
		}
	}
	return out
}

func whittle(input []string, whittler func([]string, int) uint8) string {
	size := len(input[0])
	discarded := make([]bool, len(input))
	for i := 0; i < size; i++ {
		kept := drop(input, discarded)
		if len(kept) == 1 {
			break
		}
		match := whittler(kept, i)
		for j, s := range input {
			if s[i] != match {
				discarded[j] = true
			}
		}
	}
	return drop(input, discarded)[0]
}

func commonestBitAt(ss []string, place int) uint8 {
	return extremeBitAtPlace(ss, place, func(o int, z int) uint8 {
		if o >= z {
			return '1'
		}
		return '0'
	})
}

func rarestBitAt(ss []string, place int) uint8 {
	return extremeBitAtPlace(ss, place, func(o int, z int) uint8 {
		if z <= o {
			return '0'
		}
		return '1'
	})
}

func extremeBitAtPlace(ss []string, place int, chooser func(int, int) uint8) uint8 {
	ones, zeroes := 0, 0
	for _, s := range ss {
		if s[place] == '1' {
			ones += 1
		} else {
			zeroes += 1
		}
	}
	return chooser(ones, zeroes)
}

func bitsToInt(bits string) int {
	i, _ := strconv.ParseInt(bits, 2, 64)
	return int(i)
}
