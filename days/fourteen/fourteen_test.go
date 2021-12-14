package fourteen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayFourteenPartOne(t *testing.T) {
	testCases := []struct {
		path     string
		expected int
	}{
		{path: "testdata/example", expected: 1588},
		{path: "../../inputs/fourteen", expected: 2068},
	}
	for _, tt := range testCases {
		formula := GetInput(tt.path)
		assert.Equal(t, tt.expected, PartOne(formula, 10))
	}
}
func TestDayFourteenPartTwo(t *testing.T) {
	testCases := []struct {
		path     string
		expected int
	}{
		{path: "../../inputs/fourteen", expected: 1588},
	}
	for _, tt := range testCases {
		formula := GetInput(tt.path)
		for i := 0; i < 20; i++ {
			fmt.Println(formula.Iterate())
		}
	}
}
