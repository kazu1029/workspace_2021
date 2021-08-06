package main

import "fmt"

func encrypt(str string, key int) string {
	shift, offset := rune(key%26), rune(26)
	fmt.Printf("shift=%v, offset=%v\n", shift, offset)
	runes := []rune(str)
	for i, char := range runes {
		if char >= 'a' && char+shift <= 'z' {
			char += shift
		} else {
			char += shift - offset
		}
		runes[i] = char
	}

	return string(runes)
}

func main() {
	fmt.Println("str = 'abc', key = 2, encrypted = ", encrypt("abc", 2))
	fmt.Println("str = 'opqrstu', key = 4, encrypted = ", encrypt("opqrstu", 4))
	fmt.Println("str = 'vwxyz', key = 3, encrypted = ", encrypt("vwxyz", 3))
}
