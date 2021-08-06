package main

import (
	"fmt"
	"unicode"
)

type Cipher interface {
	Encrypt(string) string
	Decrypt(string) string
}

type cipher []int

func (c cipher) cipherAlgorithm(letters string, shift func(int, int) int) string {
	shiftedText := ""
	for _, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue
		}
		shiftDist := c[len(shiftedText)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		switch {
		case s < 'a':
			s += 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}

func (c *cipher) Encrypt(str string) string {
	return c.cipherAlgorithm(str, func(a, b int) int { return a + b })
}

func (c *cipher) Decrypt(str string) string {
	return c.cipherAlgorithm(str, func(a, b int) int { return a - b })
}

func NewCaesar(key int) Cipher {
	return NewShift(key)
}

func NewShift(shift int) Cipher {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	c := cipher([]int{shift})
	return &c
}

func main() {
	c := NewCaesar(1)
	fmt.Println("Encrypt Key(01) abcd =>", c.Encrypt("abcd"))
	fmt.Println("Encrypt Key(01) bcde =>", c.Decrypt("bcde"))
	fmt.Println()

	c = NewCaesar(10)
	fmt.Println("Encrypt Key(10) abcd =>", c.Encrypt("abcd"))
	fmt.Println("Encrypt Key(10) klmn =>", c.Decrypt("klmn"))
	fmt.Println()

	c = NewCaesar(15)
	fmt.Println("Encrypt Key(15) abcd =>", c.Encrypt("abcd"))
	fmt.Println("Encrypt Key(15) pqrs =>", c.Decrypt("pqrs"))
}
