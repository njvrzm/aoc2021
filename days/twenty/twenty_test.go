package twenty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayTwenty(t *testing.T) {
	testCases := []struct {
		w World
		g int
		e int
	}{
		{
			w: GetInput("testdata/example"),
			g: 2,
			e: 35,
		},
		{
			w: GetInput("../../inputs/twenty"),
			g: 2,
			e: 5306,
		},
		{
			w: GetInput("testdata/example"),
			g: 50,
			e: 3351,
		},
		{
			w: GetInput("../../inputs/twenty"),
			g: 50,
			e: 17497,
		},
	}
	for _, tt := range testCases {
		for i := 0; i < tt.g; i++ {
			tt.w.Tick()
		}
		assert.Equal(t, tt.e, tt.w.Count())
	}
}
func TestDayTwenty_PartTwo(t *testing.T) {
}
