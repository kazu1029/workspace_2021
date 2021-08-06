package cipher

import (
	"strings"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type MyCipher []int

func (c MyCipher) code(s string, x int) string {
	ns := len(c)
	var i int
	s = strings.Map(func(r rune) rune {
		if x == 1 {
			if !unicode.IsLetter(r) {
				return -1
			}
			if unicode.IsUpper(r) {
				r += 32
			}
		}
		r += rune((c[i%ns] % 26) * x)
		if r > 'z' {
			r = 'a' + (r - ('z' + 1))
		} else if r < 'a' {
			r = 'z' - (('a' - 1) - r)
		}
		i++
		return r
	}, s)

	return s
}

func (c MyCipher) Encode(s string) string {
	return c.code(s, 1)
}

func (c MyCipher) Decode(s string) string {
	return c.code(s, -1)
}

func NewCaesar() MyCipher {
	return MyCipher{3}
}

func NewShift(n int) MyCipher {
	if n == 0 || n < -25 || n > 25 {
		return nil
	}
	return MyCipher{n}
}
