package generatedocument1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GenerateDocument(t *testing.T) {
	testCases := []struct {
		name       string
		characters string
		document   string
		expected   bool
	}{
		{
			name:       "case1",
			characters: "Bste!hetsi Expert x ",
			document:   "Expert is the Best!",
			expected:   true,
		},
		{
			name:       "case2",
			characters: "A",
			document:   "a",
			expected:   false,
		},
		{
			name:       "case3",
			characters: "a",
			document:   "a",
			expected:   true,
		},
		{
			name:       "case4",
			characters: "a hsgalhsa sanbjksbdkjba kjx",
			document:   "",
			expected:   true,
		},
		{
			name:       "caes5",
			characters: " ",
			document:   "hello",
			expected:   false,
		},
		{
			name:       "caes6",
			characters: "     ",
			document:   "     ",
			expected:   true,
		},
		{
			name:       "caes7",
			characters: "aheaollabbhb",
			document:   "hello",
			expected:   true,
		},
		{
			name:       "caes8",
			characters: "aheaolabbhb",
			document:   "hello",
			expected:   false,
		},
		{
			name:       "caes9",
			characters: "estssa",
			document:   "testing",
			expected:   false,
		},
		{
			name:       "caes10",
			characters: "Bste!hets Expert",
			document:   "Expert is the Best!",
			expected:   false,
		},
		{
			name:       "caes11",
			characters: "abcabcabcacbcdaabc",
			document:   "bacaccadac",
			expected:   true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual := GenerateDocument(test.characters, test.document)
			require.Equal(t, test.expected, actual)
		})
	}
}
