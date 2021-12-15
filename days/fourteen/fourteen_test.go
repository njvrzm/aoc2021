package fourteen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayFourteen(t *testing.T) {
	testCases := []struct {
		path       string
		iterations int
		expected   int
	}{
		{path: "testdata/example", iterations: 10, expected: 1588},
		{path: "../../inputs/fourteen", iterations: 10, expected: 2068},
		{path: "testdata/example", iterations: 40, expected: 2188189693529},
		{path: "../../inputs/fourteen", iterations: 40, expected: 2158894777814},
	}
	for _, tt := range testCases {
		formula := GetInput(tt.path)
		assert.Equal(t, tt.expected, PartOne(formula, tt.iterations))
	}
}
