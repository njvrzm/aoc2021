package sixteen

import (
	"aoc/help"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestDaySixteenPartOne(t *testing.T) {
	testCases := []struct {
		hex        string
		versionSum int
	}{
		{
			"D2FE28",
			6,
		},
		{
			"38006F45291200",
			9,
		},
		{
			"8A004A801A8002F478",
			16,
		},
		{
			"620080001611562C8802118E34",
			12,
		},
		{
			"EE00D40C823060",
			14,
		},
		{
			"C0015000016115A2E0802F182340",
			23,
		},
		{
			"A0016C880162017C3686B18A3D4780",
			31,
		},
		{
			help.ReadInput("../../inputs/sixteen")[0],
			943,
		},
	}
	for _, tt := range testCases {
		r := NewReader(tt.hex)
		p := r.Next()
		assert.Equal(t, tt.versionSum, p.VersionSum())
	}
}
func TestDaySixteenPartTwo(t *testing.T) {
	testCases := []struct {
		hex   string
		value int
	}{
		{
			"C200B40A82",
			3,
		},
		{
			"04005AC33890",
			54,
		},
		{
			"880086C3E88112",
			7,
		},
		{
			"CE00C43D881120",
			9,
		},
		{
			"D8005AC2A8F0",
			1,
		},
		{
			"F600BC2D8F",
			0,
		},
		{
			"9C005AC2F8F0",
			0,
		},
		{
			"9C0141080250320F1802104A08",
			1,
		},
		{
			help.ReadInput("../../inputs/sixteen")[0],
			167737115857,
		},
	}
	for _, tt := range testCases {
		r := NewReader(tt.hex)
		p := r.Next()
		//fmt.Println(p.Span, len(tt.hex))
		assert.Equal(t, tt.value, p.Evaluate(), tt.hex)
	}
}
