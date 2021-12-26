package eighteen

import (
	"aoc/help"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayEighteen_Misc(t *testing.T) {
	testCases := []struct {
		number string
		depth  int
	}{
		{
			"[3,[4,5]]",
			2,
		},
		{
			"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
			4,
		},
	}

	for _, tt := range testCases {
		r := Reader{string: tt.number}
		n := r.ReadNumber()
		assert.Equal(t, tt.depth, n.Depth())
		assert.Equal(t, tt.number, n.ToString())

	}
}
func TestDayEighteen_Explode(t *testing.T) {
	testCases := []struct {
		before string
		after  string
	}{
		{
			"[[[[[9,8],1],2],3],4]",
			"[[[[0,9],2],3],4]",
		},
		{
			"[7,[6,[5,[4,[3,2]]]]]",
			"[7,[6,[5,[7,0]]]]",
		},
	}

	for _, tt := range testCases {
		r := Reader{string: tt.before}
		n := r.ReadNumber()
		assert.True(t, n.Explode())
		after := n.ToString()
		assert.Equal(t, tt.after, after)
		assert.False(t, n.Explode())
	}
}

// We test explode-then-split so we don't need to adapt the
// parser to parse the invalid temporary numbers arising
// from explosion
func TestDayEighteen_PostExplodeSplit(t *testing.T) {
	testCases := []struct {
		beforeExplode string
		afterSplits   string
		splits        int
	}{
		{
			beforeExplode: "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
			afterSplits:   "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
			splits:        2,
		},
	}

	for _, tt := range testCases {
		r := Reader{string: tt.beforeExplode}
		n := r.ReadNumber()
		n.Explode()
		for count := 0; count < tt.splits; count++ {
			assert.True(t, n.Split())
		}
		assert.Equal(t, tt.afterSplits, n.ToString())
		assert.False(t, n.Split())
	}

}

func TestDayEighteen_Add(t *testing.T) {
	testCases := []struct {
		numbers   []*Number
		sum       string
		magnitude int
	}{
		{
			ReadNumbers([]string{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]"}),
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			1384,
		},
		{
			ReadNumbers([]string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			}),
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			3488,
		},
		{
			ReadNumbers(help.ReadInput("../../inputs/eighteen")),
			"[[[[6,6],[6,6]],[[6,7],[7,7]]],[[[0,7],[7,7]],[[7,7],[7,8]]]]",
			3734,
		},
	}
	for _, tt := range testCases {
		sum := Sum(tt.numbers)
		assert.Equal(t, tt.sum, sum.ToString())
		assert.Equal(t, tt.magnitude, sum.Magnitude())
	}
}

func TestDayEighteen_Magnitude(t *testing.T) {
	testCases := []struct {
		number    string
		magnitude int
	}{
		{
			"[[1,2],[[3,4],5]]",
			143,
		},
		{
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			1384,
		},
		{
			"[[[[1,1],[2,2]],[3,3]],[4,4]]",
			445,
		},
		{
			"[[[[3,0],[5,3]],[4,4]],[5,5]]",
			791,
		},
		{
			"[[[[5,0],[7,4]],[5,5]],[6,6]]",
			1137,
		},
		{
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			3488,
		},
	}

	for _, tt := range testCases {
		assert.Equal(t, tt.magnitude, NumberFromString(tt.number).Magnitude())
	}
}

func TestDayEighteen_PartTwo(t *testing.T) {
	testCases := []struct {
		numbers  []*Number
		maxitude int
	}{
		{
			ReadNumbers([]string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			}),
			3993,
		},
		{
			ReadNumbers(help.ReadInput("../../inputs/eighteen")),
			3734,
		},
	}
	for _, tt := range testCases {
		assert.Equal(t, tt.maxitude, MaximumBinarySum(tt.numbers))
	}
}
