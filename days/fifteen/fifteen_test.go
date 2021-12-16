package fifteen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayFifteen(t *testing.T) {
	testCases := []struct {
		path     string
		magnify  bool
		expected int
	}{
		{path: "testdata/example", expected: 40},
		{path: "../../inputs/fifteen", expected: 698},
		{path: "testdata/example", expected: 315, magnify: true},
		{path: "../../inputs/fifteen", expected: 3022, magnify: true},
	}
	for _, tt := range testCases {
		rm := GetInput(tt.path)
		if tt.magnify {
			rm.Magnify()
		}
		risk := rm.DialItIn()
		assert.Equal(t, tt.expected, risk)
	}
}
func TestDayFifteen_PartTwo(t *testing.T) {
}
