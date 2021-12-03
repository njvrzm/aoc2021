package three

import (
	"aoc/help"
	"strconv"
)

func GetInput(path string) []string {
	return help.ReadInput(path)
}

func PartOne(input []string) int {
	gamma, epsilon := 0, 0
	length := len(input)
	bit := 1 << (len(input[0]) - 1)
	for i := range input[0] {
		ones := 0
		for _, str := range input {
			if str[i:i+1] == "1" {
				ones += 1
			}
		}
		if ones > length / 2 {
			gamma += bit
		} else {
			epsilon += bit
		}
		bit /= 2
	}
	return gamma * epsilon
}

func PartTwo(input []string) int {
	oxy := make([]string, len(input))
	copy(oxy, input)
	for i := range oxy[0] {
		if len(oxy) == 1 {
			break
		}
		ones, zeros := 0, 0
		for _, str := range oxy {
			if str[i] == '1' {
				ones += 1
			} else {
				zeros += 1
			}
		}
		var newoxy []string
		if ones >= zeros {
			newoxy = []string{}
			for _, str := range oxy {
				if str[i] == '1' {
					newoxy = append(newoxy, str)
				}
			}
		} else {
			newoxy = []string{}
			for _, str := range oxy {
				if str[i] == '0' {
					newoxy = append(newoxy, str)
				}
			}
		}
		oxy = newoxy
	}
	co2 := make([]string, len(input))
	copy(co2, input)
	for i := range co2[0] {
		if len(co2) == 1 {
			break
		}
		ones, zeros := 0, 0
		for _, str := range co2 {
			if str[i] == '1' {
				ones += 1
			} else {
				zeros += 1
			}
		}
		newco2 := []string{}
		if zeros <= ones {
			newco2 = []string{}
			for _, str := range co2 {
				if str[i] == '0' {
					newco2 = append(newco2, str)
				}
			}
		} else {
			newco2 = []string{}
			for _, str := range co2 {
				if str[i] == '1' {
					newco2 = append(newco2, str)
				}
			}
		}
		co2 = newco2
	}

	o, _ := strconv.ParseInt(oxy[0], 2, 16)
	c, _ := strconv.ParseInt(co2[0], 2, 16)
	return int(o * c)

}