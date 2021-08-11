package spiraltraverse1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpiralTraverse(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		expected []int
	}{
		{
			matrix: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			matrix: [][]int{
				{1},
			},
			expected: []int{1},
		},
		{
			matrix: [][]int{
				{1, 2},
				{4, 3},
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			matrix: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			actual := SpiralTraverse(tc.matrix)
			require.Equal(t, tc.expected, actual)
		})
	}
}
