package runlengthencoding1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := "AAAAAAAAAAAAABBCCCCDD"
	expected := "9A4A2B4C2D"
	actual := RunLengthEncode(input)
	require.Equal(t, expected, actual)
}
